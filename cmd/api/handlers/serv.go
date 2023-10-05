package handlers

import (
	"del/service"
	"net/http"
)

type IHandlers interface {
	AddOrder(w http.ResponseWriter, r *http.Request)
	GetOrder(w http.ResponseWriter, r *http.Request)
}

func NewHandler(serv service.IService) IHandlers {
	return &handler{
		serv: serv,
	}
}

type handler struct {
	serv service.IService
}
