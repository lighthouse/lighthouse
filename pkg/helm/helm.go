package helm

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
)

var chartsGithubURL = "https://github.com/kubernetes/charts"
var chartsDirName = "helm_charts"
var templatesDirName = "templates"

func downloadHelmChartsRepository(outputDir string) error {
	_, err := git.PlainClone(outputDir, false, &git.CloneOptions{
		URL:      chartsGithubURL,
		Progress: os.Stdout,
	})

	if err != nil {
		return err
	}

	return nil
}

func GetYamlsForChart(name string) ([][]byte, error) {
	yamls := [][]byte{}

	prefixDir := name[0:strings.Index(name, "/")]
	chartDir := name[strings.Index(name, "/")+1:]

	dir, _ := ioutil.TempDir("", chartsDirName)
	err := downloadHelmChartsRepository(dir)
	if err != nil {
		return nil, err
	}

	defer os.RemoveAll(dir)

	templatesDir := path.Join(path.Join(path.Join(dir, prefixDir), chartDir), templatesDirName)
	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		fileName := f.Name()
		if strings.Contains(fileName, ".yaml") {
			data, err := ioutil.ReadFile(path.Join(templatesDir, fileName))
			if err != nil {
				return nil, err
			}

			yamls = append(yamls, data)
		}
	}

	return yamls, nil
}
