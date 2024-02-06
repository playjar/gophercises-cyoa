package main

import (
	"flag"

	"github.com/playjar/gophercises-cyoa/internal/http"
	"github.com/playjar/gophercises-cyoa/internal/util"
)

func main() {
	arcJSONFilePath := flag.String("arc-json", "", "JSON file containing all arcs and their relationship")
	arcTemplateFilePath := flag.String("arc-template", "", "HTML template file for rendering arcs")
	flag.Parse()

	arcs, err := util.ParseJSONToArcs(*arcJSONFilePath)
	if err != nil {
		panic(err)
	}

	arcTemplate, err := util.ParseArcTemplate(*arcTemplateFilePath)
	if err != nil {
		panic(err)
	}

	mux := http.GetArcHandlerMux(arcs, arcTemplate)
	http.StartHTTPServerWithHandler(mux)
}