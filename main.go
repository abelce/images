package main

import (
	"admin/application"
	// "encoding/json"
	// "fmt"
	"log"
	"net/http"
)

type contentTypeMiddleware struct {
	next http.Handler
}

func (h *contentTypeMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Header().Set("access-control-allow-origin", "*")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return;
	}
	//判断cookie是否存在
	// cookie, err := r.Cookie("gosessionid")
	// if r.URL.Path != "/login" && (err != nil || cookie.Value == "") {
	// 	message := struct {
	// 		Message string
	// 	}{
	// 		Message: "session is expiresed",
	// 	}

	// 	result, _ := json.Marshal(message)
	// 	wc := 0

	// 	for wc < len(result) {
	// 		n, err := w.Write(result)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		wc += n
	// 	}
	// } else {
		h.next.ServeHTTP(w, r)
	// }
}

func main() {

	var routeHandler http.Handler = &contentTypeMiddleware{
		next: application.NewRouter(),
	}

	err := http.ListenAndServe(":9001", routeHandler)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
