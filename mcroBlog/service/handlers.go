package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go_mod-project/mcroBlog/dbclient"
	"net/http"
	"strconv"
)

var DBClient dbclient.IBotClient

func GetAccount(w http.ResponseWriter, r *http.Request)  {
	var accountId = mux.Vars(r)["accountId"]
	account, err := DBClient.QueryAccount(accountId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}