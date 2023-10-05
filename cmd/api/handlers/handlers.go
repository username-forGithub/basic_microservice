package handlers

import (
	"del/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *handler) AddOrder(w http.ResponseWriter, r *http.Request) {
	var order *models.Order

	getBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err, "getBody")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err, "Body.Close()")
		}
	}(r.Body)
	err = json.Unmarshal(getBody, &order)
	if err != nil {
		log.Println(err, "json.Unmarshal")
	}
	resp := h.serv.AddOrder(order)

	w.WriteHeader(http.StatusCreated)

	_, err = w.Write([]byte(strconv.Itoa(resp)))
	if err != nil {
		log.Println(err, "strconv.Itoa(resp))")
	}
	//clientp.Clientf()
}

func (h *handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["id"]
	orId, err := strconv.Atoi(orderId)
	if err != nil {
		log.Println(err)
	}
	orderInfoFromDb := h.serv.GetOrder(orId)
	jsonData, err := json.Marshal(orderInfoFromDb)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonData)
	if err != nil {
		log.Println(err, "writing response")
		return
	}
}
