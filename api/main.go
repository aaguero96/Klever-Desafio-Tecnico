package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aaguero96/Klever-Desafio-Tecnico/api/router"
	"github.com/aaguero96/Klever-Desafio-Tecnico/config"
)

func main() {
	config.LoadEnv()

	fmt.Println("Rodando API")

	router := router.Router()

	fmt.Printf("Rodando na porta %d\n", config.Port)
	address := fmt.Sprintf(":%d", config.Port)
	log.Fatal(http.ListenAndServe(address, router))
}
