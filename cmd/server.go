package cmd

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lighthouse/lighthouse/pkg/entities"
	"github.com/lighthouse/lighthouse/pkg/helm"
)

var dir = "web/dist"

type Chart struct {
	Name string `json:"name"`
}

func RunServer() {
	r := mux.NewRouter()

	r.HandleFunc("/chart", ProcessHelmChart).Methods("POST")
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	log.Fatal(http.ListenAndServe(":8000", r))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ProcessHelmChart(w http.ResponseWriter, r *http.Request) {
	var chart Chart
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&chart); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	yamls, err := helm.GetYamlsForChart(chart.Name)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	objs := entities.GetKubernetesObjects(yamls)
	respondWithJSON(w, http.StatusOK, objs)
}
