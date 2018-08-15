package application

import (
	"admin/domain"
	"admin/port"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"fmt"
	_ "admin/memory"
	"admin/session"
	"time"
)

var globalSessions *session.Manager

type Result struct {
	TextStatus string      `json:"textStatus"`
	Data       interface{} `json:"data"`
	Error      error       `json:"error"`
	Mate       domain.Mate  `json:"mate"`
}

func NewRouter() *mux.Router {
	fmt.Println("routers")
	r := mux.NewRouter()
	r.HandleFunc("/image", SaveImage).Methods(http.MethodPost)
	// r.HandleFunc("/article/list", ArticleList).Methods(http.MethodGet)
	// r.HandleFunc("/article/{id}", UpdateArticle).Methods(http.MethodPut)
	// r.HandleFunc("/article/{id}", GetArticle).Methods(http.MethodGet)
	// r.HandleFunc("/article/{id}", DeleteArticle).Methods(http.MethodDelete)

	return r
}