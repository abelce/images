package application

import (
	"admin/domain"
	"admin/port"
	// "github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	// "strings"
	// "strconv"
	"github.com/satori/go.uuid"
	"fmt"
	"time"
)

func SaveArticle(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm();
	r.ParseForm()
	data, err := ioutil.ReadAll(r.Body);
	if (err != nil ) {
		log.Fatal("获取参数失败")
		return;
	}
	
	article := domain.Article{}
	fmt.Println("123")
	err = json.Unmarshal(data, &article)
	fmt.Println(article.Title);		
	
	if err != nil {
		log.Fatal("解析失败")
		return;
	}

	id := article.ID;
	
	time := time.Now().Unix();
	if id == "" {
		id, _ := uuid.NewV4()
		article.ID = id.String();
		article.CreateTime = time;
	}
	article.LastUpdateTime = time;

	fmt.Println(article.Title);	
	tmp, err := port.CreateArticle(&article);
	fmt.Println(article.Title);

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