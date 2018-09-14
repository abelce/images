package main

import (
	"fmt"
	"log"
	"net/http"

	"images/application"
)

type contentTypeMiddleware struct {
	next http.Handler
}

func (h *contentTypeMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Header().Set("access-control-allow-origin", "*")
	w.Header().Set("access-control-allow-methods", "*")
	w.Header().Set("access-control-allow-headers", "*")
	w.Header().Set("access-control-expose-headers", "*")
	w.Header().Set("access-control-allow-credentials", "true")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	h.next.ServeHTTP(w, r)
}

func main() {

	path := "/go/src/images/config.json"
	cxt, err := application.NewContext(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	application.ApplicationContext = cxt

	var routeHandler http.Handler = &contentTypeMiddleware{
		next: application.NewRouter(),
	}

	err = http.ListenAndServe(":9012", routeHandler)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
