package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/rachmatprabowo/redsfin2/core"
	"github.com/rachmatprabowo/redsfin2/modules/auth/model"
)

// UserHandler is handle function
func UserHandler(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)

	switch method := r.Method; method {
	case "GET":
		if params["id"] != "" {
			id, _ := strconv.Atoi(params["id"])
			var user model.User
			if user.Load(id) {
				payload, err := json.Marshal(user)
				core.CheckErr(err, "")
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(payload)
			}
		} else {
			var users model.Users
			if users.LoadAll() {
				payload, err := json.Marshal(users)
				core.CheckErr(err, "")
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(payload)
			}
		}

	case "POST":
		if params["id"] == "" {
			decoder := json.NewDecoder(r.Body)
			var usr model.User
			err := decoder.Decode(&usr)
			_ = core.CheckErr(err, "unkown error was occured")
			r.Body.Close()
			if usr.Save() {
				w.WriteHeader(http.StatusCreated)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	}
}
