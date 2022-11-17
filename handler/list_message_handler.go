package handler

import (
	"encoding/json"
	"github.com/FlowerBirds/go-server-admin/dao"
	"github.com/FlowerBirds/go-server-admin/model"
	"log"
	"net/http"
	"strconv"
)

func MakeListMessageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		r.ParseForm()
		currentPage, err := strconv.Atoi(r.FormValue("currentPage"))
		if err != nil {
			currentPage = 1
		}
		pageSize, err := strconv.Atoi(r.FormValue("pageSize"))
		if err != nil {
			pageSize = 10
		}
		currentClient := r.FormValue("currentClient")
		messages := dao.ListClientMessages(currentPage, pageSize, currentClient)
		total := dao.CountClientMessages(currentClient)

		data := new(model.ClientMessageVO)
		data.Total = total
		data.Data = messages
		data.CurrentPage = currentPage
		data.PageSize = pageSize
		msg, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(msg)

	}
}
