package application

import (
	"admin/domain"
	"admin/port"
	"github.com/gorilla/mux"
	"net/http"
	// "io/ioutil"
	"log"
	"encoding/json"
	// "strings"
	// "strconv"
	// "github.com/satori/go.uuid"
	// "time"
	// "github.com/gorilla/mux"
)

func GetTypesByType(w http.ResponseWriter, r *http.Request) {
	r.ParseForm();

	vars := mux.Vars(r);
	_type := vars["type"]

	if (_type == "") {
		log.Fatal("type is undefined")
		return
	}

	tmp, err := port.GetTypeList(_type)
	res := Result{}

	if err != nil {
		res.TextStatus = "failed"
		res.Error = err
	} else {
		res.TextStatus = "ok"
		res.Data = tmp
		mate := domain.Mate{}
		mate.Total = len(tmp)
		res.Mate = mate
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