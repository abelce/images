package application

import (
	"admin/domain"
	"admin/port"
	// "github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	// "log"
	"encoding/json"
	// "strings"
	"strconv"
	"github.com/satori/go.uuid"
	"time"
	"github.com/gorilla/mux"
	"admin/utils"
	"errors"
)

func GetTimestampString() string {
	return strconv.FormatInt(time.Now().Unix(), 10);
}

func SaveArticle(w http.ResponseWriter, r *http.Request) {
	utils.ResetHTTPErrors()
	r.ParseForm()
	data, err := ioutil.ReadAll(r.Body);
	if err != nil {
		utils.HandleHTTPError(w, err)
	}
	article := domain.Article{}
	err = json.Unmarshal(data, &article)
	if err != nil {
		utils.HandleHTTPError(w, err)
	}
	id := article.ID;
	time := GetTimestampString();
	if id == "" {
		id, _ := uuid.NewV4()
		article.ID = id.String();
		article.CreateTime = time;
	}
	article.LastUpdateTime = time;
	tmp, err := port.CreateArticle(&article);
	if err != nil {
		utils.HandleHTTPError(w, err);
	}
	res := Result{
		TextStatus: "ok",
		Data: tmp,
	};
	result, _ := json.Marshal(res)
	wc := 0
	for wc < len(result) {
		n, err := w.Write(result)
		if err != nil {
			utils.HandleServerError(w, err)
		}
		wc += n
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
	article := domain.Article{}
	err = json.Unmarshal(data, &article)
	if err != nil {
		utils.HandleHTTPError(w, err)
	}
	//更新LastUpdateTime
	time := GetTimestampString();
	article.LastUpdateTime = time;

	tmp, err := port.UpdateArticle(id, &article);
	if err != nil {
		utils.HandleHTTPError(w, err)
	}
	res := Result{
		TextStatus: "ok",
		Data: tmp,
	};
	result, _ := json.Marshal(res)
	wc := 0
	for wc < len(result) {
		n, err := w.Write(result)
		if err != nil {
			utils.HandleServerError(w, err)
		}
		wc += n
	}
}


func GetArticle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.HandleHTTPError(w, errors.New("id is null"))
	}
	tmp, err := port.GetArticle(id)
	if err != nil {
		utils.HandleHTTPError(w, err)
	}
	res := Result{}
	res.TextStatus = "ok"
	res.Data = tmp
	result, _ := json.Marshal(res)
	wc := 0
	for wc < len(result) {
		n, err := w.Write(result)
		if err != nil {
			utils.HandleServerError(w, err)
		}
		wc += n
	}
}


func ArticleList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var pageSize int = 10
	var pageNum int = 1
	if r.FormValue("pageSize") != "" {
		pageSize, _ = strconv.Atoi(r.FormValue("pageSize"))
	}
	if r.FormValue("pageNum") != "" {
		pageNum, _ = strconv.Atoi(r.FormValue("pageNum"))
	}

	tmp, err := port.GetArticleList(pageSize * (pageNum - 1), pageNum * pageSize )
	if err != nil {
		utils.HandleHTTPError(w, err)
	}
	res := Result{}
	res.TextStatus = "ok"
	res.Data = tmp

	mate := domain.Mate{}
	total, _ := port.ArticleTotal();
	mate.Total = total
	res.Mate = mate

	result, _ := json.Marshal(res)

	wc := 0
	for wc < len(result) {
		n, err := w.Write(result)
		if err != nil {
			// panic(err)
			utils.HandleServerError(w, err)
		}
		wc += n
	}
}