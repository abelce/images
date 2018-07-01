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
	// "strconv"
	// "time"
	// "github.com/gorilla/mux"
	"admin/utils"
	"github.com/satori/go.uuid"
	// "fmt"
)

func GetTypesByType(w http.ResponseWriter, r *http.Request) {
	r.ParseForm();

	tmp, err := port.GetTypeList()
	if err != nil {
		utils.HandleServerError(w, err)
	}
	res := Result{
		TextStatus: "ok",
		Data: tmp,
	}
	
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


func CreateType(w http.ResponseWriter, r *http.Request) {
	utils.ResetHTTPErrors()
	r.ParseForm();
	data, err := ioutil.ReadAll(r.Body);
	if err != nil {
		utils.HandleHTTPError(w, err);
	}
	id, _ := uuid.NewV4();
	newType := domain.Type{}
	err = json.Unmarshal(data, &newType)
	if err != nil {
		utils.HandleHTTPError(w, err)
	}
	newType.ID = id.String()
	tmp, err := port.CreateType(&newType);
	if err != nil {
		utils.HandleHTTPError(w, err)
	}
	res := Result{
		TextStatus: "ok",
		Data: tmp,
	}
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