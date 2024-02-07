package http

import (
	"fmt"
	"html/template"
	net_http "net/http"

	"github.com/playjar/gophercises-cyoa/internal/model"
)

type arcHandler struct {
	CurrentArc model.Arc
	Template *template.Template
}

func NewArcHandler(arc model.Arc, tmpl *template.Template) net_http.Handler {
	return arcHandler{CurrentArc: arc, Template: tmpl}
}

func (ah arcHandler) ServeHTTP(w net_http.ResponseWriter, r *net_http.Request) {
	err := ah.Template.Execute(w, ah.CurrentArc)
	if err != nil {
		fmt.Printf("arc template execution error: %v, request: %+v\n", err, r)
	}
}