package application

import (
	"bytes"
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	// "strings"
	"strconv"
	"time"
	"github.com/gorilla/mux"
	"admin/utils"
	"errors"
	"github.com/google/jsonapi"
	"admin/application/command"
	sjson "github.com/bitly/go-simplejson"
)

const successJSON = `{
	"responseStatus": {
		"success": true,
		"version": "v1"
	}
}`

func GetTimestampString() string {
	return strconv.FormatInt(time.Now().Unix(), 10);
}

func SaveArticle(w http.ResponseWriter, r *http.Request) {
	utils.ResetHTTPErrors()
	r.ParseForm()
	data, err := ioutil.ReadAll(r.Body);
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}
	a := new(command.CreateArticle)
	err = json.Unmarshal(data, &a)
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}

	service, err := ApplicationContext.Service()
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}

	article, err :=  service.CreateArticle(*a)
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Println(article.ID)
	err = jsonapi.MarshalPayloadWithoutIncluded(w, article)
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.HandleHTTPError(w, errors.New("id is null"))
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.HandleHTTPError(w, err)
	}
	a := new(command.UpdateArticle)
	err = json.Unmarshal(data, &a)
	if err != nil {
		utils.HandleHTTPError(w, err)
	}
	a.ID = id
	service, err := ApplicationContext.Service()
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}
	article, err := service.UpdateArticle(*a)
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = jsonapi.MarshalPayloadWithoutIncluded(w, article)
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}
}


func GetArticle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.HandleHTTPError(w, errors.New("id is null"))
	}
	qs, err := ApplicationContext.QueryService()
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}
	article, err := qs.FindByID(id)
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}
	err = jsonapi.MarshalPayloadWithoutIncluded(w, article)
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}
}


func ArticleList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var pageSize int = 100
	var pageNum int = 1
	if r.FormValue("pageSize") != "" {
		pageSize, _ = strconv.Atoi(r.FormValue("pageSize"))
	}
	if r.FormValue("pageNum") != "" {
		pageNum, _ = strconv.Atoi(r.FormValue("pageNum"))
	}

	offsetNum := (pageNum - 1) * pageSize
	limitNum := pageNum * pageSize
	qs, err := ApplicationContext.QueryService()
	total, articles, err := qs.Find(offsetNum, limitNum)
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}
	trw := bytes.NewBuffer(nil)
	err = jsonapi.MarshalPayloadWithoutIncluded(trw, articles)
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}
	json, err := sjson.NewFromReader(trw)
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}
	json.SetPath([]string{"mate", "total"}, total)
	data, err := json.Encode()
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}
	w.Write(data)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id := mux.Vars(r)["id"]
	c := command.DeleteArticle{
		ID: id,
	}
	service, err := ApplicationContext.Service()
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}
	err = service.DeleteArticle(c)
	if err != nil {
		utils.HandleHTTPError(w, err)
		return
	}

	w.Write([]byte(successJSON))
}