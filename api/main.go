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

	fmt.Printf("Rodando na porta %d\n", 5000)
	address := fmt.Sprintf(":%d", 5000)
	log.Fatal(http.ListenAndServe(address, router))
}
