package entities

import (
	"bufio"
	"strings"
)

type KubernetesRootObject struct {
	Kind      string                   `json:"kind"`
	Metadata  KubernetesMetadataObject `json:"metadata"`
	RawString string                   `json:"rawString"`
}

type KubernetesMetadataObject struct {
	Name string `json:"name"`
}

func GetKubernetesObjects(templates [][]byte) []KubernetesRootObject {
	kubernetesObjects := []KubernetesRootObject{}

	for _, t := range templates {
		rawString := string(t)
		scanner := bufio.NewScanner(strings.NewReader(rawString))
		kind := ""
		name := ""
		metadataFound := false

		for scanner.Scan() {
			line := scanner.Text()

			if metadataFound {
				if strings.Contains(line, "name:") {
					name = line[strings.Index(line, ":")+2:]
				}
				metadataFound = false
			} else if strings.Contains(line, "kind:") {
				kind = line[strings.Index(line, ":")+2:]
			} else if strings.Contains(line, "metadata:") {
				metadataFound = true
			}

			if name != "" && kind != "" {
				r := KubernetesRootObject{
					Kind: kind,
					Metadata: KubernetesMetadataObject{
						Name: name,
					},
					RawString: rawString,
				}

				kubernetesObjects = append(kubernetesObjects, r)
				name = ""
				kind = ""
			}
		}
	}

	return kubernetesObjects
}
