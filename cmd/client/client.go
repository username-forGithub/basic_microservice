package clientp

import (
	"del/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Clientf() {
	baseUrl := "http://localhost:8080/deliver/v1/get-order/6"
	client := http.Client{}
	r, err := client.Get(baseUrl)
	if err != nil {
		log.Println(err, "client.Get(baseUrl)")
	}
	defer r.Body.Close()
	respBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err, "client ReadAll(r.Body)")
	}
	orderInstance := &models.Order{}
	err = json.Unmarshal(respBody, orderInstance)
	if err != nil {
		log.Println(err, "client json.Unmarshal(respBody, orderInstance)")
	}

	fmt.Println("%+v", orderInstance.Name)
}
