package main

import (
	"net/http"
	"admin/application"
	"log"
	"encoding/json"
	// "fmt"
)

type contentTypeMiddleware struct {
	next http.Handler
}

func (h *contentTypeMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Header().Set("access-control-allow-origin","*")
	
	//判断cookie是否存在
	cookie, err := r.Cookie("gosessionid");
	if r.URL.Path !="/login" &&  (err !=nil || cookie.Value == "") {
		message := struct {
			Message string
		} {
			Message: "session is expiresed",
		}

		result, _ := json.Marshal(message)
		wc := 0

		for wc < len(result) {
			n, err := w.Write(result);
			if err != nil {
				panic(err)
			}
			wc += n;
		}
	} else {
		h.next.ServeHTTP(w, r)
	}
}

func main() {

	var routeHandler http.Handler = &contentTypeMiddleware{
		next: application.NewRouter(),
	}

	err := http.ListenAndServe(":9090", routeHandler)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}