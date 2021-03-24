package handler

import (
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type pathUrl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

// MapHandler returns a handler func
// that maps every key in the map to its value
// if the path is not provided, the fallback will be called
func MapHandler(paths map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//if path match - redirect to the url
		path := r.URL.Path
		if dest, ok := paths[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
	}
}

// YamlHandler parses the yaml and returns a handler func
// that maps every key in the map to its value
// if the path is not provided, the fallback will be called
func YamlHandler(yamlBytes []byte, fallback http.Handler) (http.Handler, error) {
	// parse yaml
	pathUrls, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}

	//convert yaml array into map key(path) - val(url)
	pathsMap := convertToMap(pathUrls)
	return MapHandler(pathsMap, fallback), nil
}

func parseYaml(yamlBytes []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yamlBytes, &pathUrls)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}
	return pathUrls, nil
}

func convertToMap(pathUrls []pathUrl) map[string]string {
	pathsMap := make(map[string]string)
	for _, v := range pathUrls {
		pathsMap[v.Path] = v.Url
	}
	return pathsMap
}
