package application

import (
	"admin/domain"
	"admin/port"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	// "strings"
	"strconv"
	"fmt"
	_ "admin/memory"
	"admin/session"
	"time"
	// "admin/application"
)

var globalSessions *session.Manager

type Result struct {
	TextStatus string      `json:"textStatus"`
	Data       interface{} `json:"data"`
	Error      error       `json:"error"`
	Mate       domain.Mate  `json:"mate"`
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {

	}

	user := domain.User{}
	json.Unmarshal(data, &user)

	email := user.Email
	passwd := user.Password
	if email == "" || passwd == "" {
		log.Fatal("用户名或密码为空")
		return
	}

	fmt.Println(email)
	newUser, err := port.Login(email, passwd)

	res := Result{}
	if err != nil {
		res.TextStatus = "failed"
		res.Error = err
	} else {
		res.TextStatus = "ok"
		res.Data = &newUser

		//设置session
		globalSessions.SessionStart(w, r)
	}

	result, _ := json.Marshal(res)

	wc := 0
	for wc < len(result) {
		n, err := w.Write(result)
		if err != nil {
			panic(err)
		}
		wc += n
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("读取用户信息创建信息错误")
		return
	}
	newUser := domain.User{}
	err = json.Unmarshal(data, &newUser)
	if err != nil {
		log.Fatal("解析失败")
		return
	}

	time := time.Now().Unix()
	newUser.AccessTime = time
	newUser.LastUpdateTime = time
	newUser.CreateTime = time

	tmp, err := port.SaveUser(&newUser)

	res := Result{}
	if err != nil {
		res.TextStatus = "failed"
		res.Error = err
	} else {
		res.TextStatus = "ok"
		res.Data = tmp
	}

	// fmt.Println(res.result)

	result, _ := json.Marshal(res)

	wc := 0
	for wc < len(result) {
		n, err := w.Write(result)
		if err != nil {
			panic(err)
		}
		wc += n
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userID, _ := ioutil.ReadAll(r.Body)

	if len(userID) == 0 {
		return
	}

	params := struct {
		Id int `json:"id"`
	}{}

	json.Unmarshal(userID, &params)

	err := port.Delete(params.Id)
	result := Result{}
	if err != nil {
		result.TextStatus = "failed"
	} else {
		result.TextStatus = "ok"
	}
	res, _ := json.Marshal(result)

	wc := 0
	for wc < len(res) {
		n, err := w.Write(res)
		if err != nil {
			panic(err)
		}
		wc += n
	}
}

func findUsers(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var pageSize int
	var pageNum int
	if r.FormValue("pageSize") != "" {
		pageSize, _ = strconv.Atoi(r.FormValue("pageSize"))
	} else {
		pageSize = 10
	}

	if r.FormValue("pageNum") != "" {
		pageNum, _ = strconv.Atoi(r.FormValue("pageNum"))
	} else {
		pageNum = 1
	}

	offset := (pageNum - 1) * pageSize
	end := pageNum * pageSize

	users, err := port.Users(offset, end)

	res := Result{}
	if err != nil {
		res.TextStatus = "failed"
		res.Error = err
	} else {
		res.TextStatus = "ok"
		res.Data = users
	}

	result, err := json.Marshal(users)

	wc := 0
	for wc < len(result) {
		n, err := w.Write(result)
		if err != nil {
			panic("查询所有用户返回写入失败")
		}
		wc += n
	}
}

// func test(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello"))
// }

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/login", login).Methods(http.MethodPost)
	r.HandleFunc("/create", create).Methods(http.MethodPost)
	r.HandleFunc("/delete", delete).Methods(http.MethodPut)
	r.HandleFunc("/users", findUsers).Methods(http.MethodGet)

	r.HandleFunc("/article", SaveArticle).Methods(http.MethodPost)
	r.HandleFunc("/article/{id}", UpdateArticle).Methods(http.MethodPut)
	r.HandleFunc("/article/{id}", GetArticle).Methods(http.MethodGet)
	r.HandleFunc("/article/list", ArticleList).Methods(http.MethodGet)

	return r
}

func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3306)
	go globalSessions.GC()
}
