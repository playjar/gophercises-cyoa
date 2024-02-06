package util

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"

	"github.com/playjar/gophercises-cyoa/internal/model"
)

func ParseJSONToArcs(filepath string) (arcs []model.Arc, err error) {
	jsonBytes, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("parse arc json file error: %v, file path: %s", err, filepath)
		return nil, err
	}
	
	var arcNameToArc map[string]model.Arc
	err = json.Unmarshal(jsonBytes, &arcNameToArc)
	if err != nil {
		fmt.Printf("unmarshal arc json bytes error: %v", err)
		return nil, err
	}

	for arcName, arc := range(arcNameToArc) {
		arc.Name = arcName
		arcs = append(arcs, arc)
	}

	return arcs, nil
}

func ParseArcTemplate(filepath string) (arcTemplate *template.Template, err error) {
	arcTemplate, err = template.ParseFiles(filepath)
	if err != nil {
		fmt.Printf("parse arc html template error: %v, file path: %s", err, filepath)
		return nil, err
	}

	return arcTemplate, nil
}