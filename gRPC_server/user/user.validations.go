package user_server

import (
	"errors"
	"log"
	"net/mail"
)

func validateName(name string) error {
	if name == "" {
		log.Println("Name is required")
		return errors.New("Name is required")
	}
	return nil
}

func validateEmail(email string) error {
	if email == "" {
		log.Println("Email is required")
		return errors.New("Email is required")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
