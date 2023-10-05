package main

import (
	"del/cmd/api/handlers"
	"del/service"
	"del/storage"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	httpPort = ":8080"
)

func main() {
	db := storage.ConnectToDb()
	getNewStorage := storage.NewStorage(db)
	getNewService := service.NewService(getNewStorage)
	handler := handlers.NewHandler(getNewService)

	router := mux.NewRouter()
	router.HandleFunc("/deliver/v1/add-order", handler.AddOrder).Methods(http.MethodPost)
	router.HandleFunc("/deliver/v1/get-order/{id}", handler.GetOrder).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:        httpPort,
		ReadTimeout: time.Second * 10,
		Handler:     router,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err, "listenAndServe")
	}

}
