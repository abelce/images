package application

import (
	"admin/port"
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	// "strings"
	"strconv"
	"github.com/satori/go.uuid"
	"fmt"
	"time"
)

func SaveArticle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm();
	data, err := ioutil.ReadAll(r.Body);
	article := port.Article{}
	err = json.Unmarshal(data, &article)
	if err != nil {
		log.Fatal("解析失败")
		return;
	}

	id := article.ID;
	fmt.Println(id)
	time := time.Now().Unix();	
	if id == nil {
		id, _ := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
		article.ID = id;
		article.CreateTime = time;
	}
	article.LastUpdateTime = time;

	tmp, err := port.CreateArticle(&article);

	res := Result{};
	if err != nil {
		res.TextStatus = "failed";
		res.Error = err;
	} else {
		res.TextStatus = "ok";
		res.Data = tmp
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