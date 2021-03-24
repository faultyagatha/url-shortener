package main

import (
	"fmt"
	"net/http"

	"github.com/faultyagatha/urlshortener/handler"
)

func main() {
	//mux is a fallback
	mux := defaultMux()

	paths := map[string]string{
		"/agatha":    "https://faultyagatha.github.io",
		"/portfolio": "https://faultyagatha.github.io/devprofile/",
	}
	mapHandler := handler.MapHandler(paths, mux)

	yaml := `
- path: /agatha
  url: https://faultyagatha.github.io
- path: /portfolio
  url: https://faultyagatha.github.io/devprofile
`

	yamlHandler, err := handler.YamlHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
