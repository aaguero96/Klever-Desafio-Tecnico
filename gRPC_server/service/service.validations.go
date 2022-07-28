package service_server

import (
	"errors"
	"log"
)

func validateName(name string) error {
	if name == "" {
		log.Println("Name is required")
		return errors.New("Name is required")
	}
	return nil
}
