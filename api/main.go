package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aaguero96/Klever-Desafio-Tecnico/api/router"
)

func main() {
	fmt.Println("Rodando API")

	router := router.Router()

	fmt.Println("Rodando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
