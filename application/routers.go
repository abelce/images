package application

import (
	"images/domain"
	// "images/port"
	// "encoding/json"
	"github.com/gorilla/mux"
	// "io/ioutil"
	// "log"
	"net/http"
	// "strconv"
	// "fmt"
	// "time"
)

type Result struct {
	TextStatus string      `json:"textStatus"`
	Data       interface{} `json:"data"`
	Error      error       `json:"error"`
	Mate       domain.Mate  `json:"mate"`
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/image", SaveImage).Methods(http.MethodPost)
	r.HandleFunc("/image", ImageList).Methods(http.MethodGet)
	// r.HandleFunc("/article/{id}", UpdateArticle).Methods(http.MethodPut)
	// r.HandleFunc("/article/{id}", GetArticle).Methods(http.MethodGet)
	// r.HandleFunc("/article/{id}", DeleteArticle).Methods(http.MethodDelete)

	return r
}