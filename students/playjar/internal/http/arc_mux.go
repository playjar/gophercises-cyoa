package http

import (
	"html/template"
	net_http "net/http"

	"github.com/playjar/gophercises-cyoa/internal/model"
)

func GetArcHandlerMux(arcs []model.Arc, arcTemplate *template.Template) (mux *net_http.ServeMux) {
	mux = net_http.NewServeMux()
	
	for _, arc := range arcs {
		var path = "/"
		if arc.Name != "intro" {
			path += arc.Name
		}
		mux.Handle(path, NewArcHandler(arc, arcTemplate))
	}

	return mux
}