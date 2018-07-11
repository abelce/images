package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"

	"admin/application"
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

	path := "./config.json"
	cxt, err := application.NewContext(path)
	if err != nil {
		fmt.Println("启动失败")
		fmt.Println(err)
		return
	}
	application.ApplicationContext = cxt

	var routeHandler http.Handler = &contentTypeMiddleware{
		next: application.NewRouter(),
	}

	err = http.ListenAndServe(":9001", routeHandler)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
