package http

import (
	"fmt"
	net_http "net/http"
)

func StartHTTPServerWithHandler(handler net_http.Handler) {
	fmt.Println("start http server at port 8080!")
	net_http.ListenAndServe(":8080", handler)
}