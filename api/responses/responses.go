package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		if erro := json.NewEncoder(w).Encode(data); erro != nil {
			log.Fatal(erro)
		}
	}
}

func Error(w http.ResponseWriter, status int, erro error) {
	JSON(w, status, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
