package user_server

import (
	"errors"
	"fmt"
	"log"
	"net/mail"
	"unicode"
)

var (
	MIN_PASSWORD = 6
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

func passwordRules(password string) error {
	var (
		hasMinimumLen = false
		hasUpper      = false
		hasLower      = false
		hasNumber     = false
		hasSpecial    = false
	)
	if len(password) >= MIN_PASSWORD {
		hasMinimumLen = true
	}
	for _, character := range password {
		switch {
		case unicode.IsUpper(character):
			hasUpper = true
		case unicode.IsLower(character):
			hasLower = true
		case unicode.IsNumber(character):
			hasNumber = true
		case unicode.IsPunct(character):
			hasSpecial = true
		case unicode.IsSymbol(character):
			hasSpecial = true
		}
	}

	validate := hasMinimumLen &&
		hasUpper &&
		hasLower &&
		hasNumber &&
		hasSpecial

	if !validate {
		text := fmt.Sprintf(`
			Password needs:
				minimun of %d chars;
				1 uppercase char;
				1 lowercase char;
				1 number char;
				1 special char.
			You forgot one or more.
		`, MIN_PASSWORD)
		log.Println(text)
		return errors.New(text)
	}

	return nil
}

func validatePassword(password string) error {
	if password == "" {
		log.Println("Password is required")
		return errors.New("Password is required")
	}

	if err := passwordRules(password); err != nil {
		return err
	}

	return nil
}
