package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yogamuris/sohappytocyou/entity"
	"image/color"
	"log"
	"net/http"
)

func UserPageHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]

	if r.Method == "GET" {
		pagePayload := entity.Page{
			Username:    username,
			Background:  color.Black,
			Photo:       "",
			Description: "Testing",
			Links:       nil,
		}
		js, err := json.Marshal(pagePayload)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(js)
	}

}