package http

import (
	"fmt"
	"html/template"
	net_http "net/http"

	"github.com/playjar/gophercises-cyoa/internal/model"
)

type ArcHandler struct {
	CurrentArc model.Arc
	Template *template.Template
}

func (ah ArcHandler) ServeHTTP(w net_http.ResponseWriter, r *net_http.Request) {
	err := ah.Template.Execute(w, ah.CurrentArc)
	if err != nil {
		fmt.Printf("arc template execution error: %v, request: %+v\n", err, r)
	}
}