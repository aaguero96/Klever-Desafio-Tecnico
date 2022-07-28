package service_server

import (
	"errors"
	"log"
	"net/url"
)

func validateName(name string) error {
	if name == "" {
		log.Println("Name is required")
		return errors.New("Name is required")
	}
	return nil
}

func validateSite(site string) error {
	if site == "" {
		log.Println("Site is required")
		return errors.New("Site is required")
	}

	response, err := url.Parse(site)
	if err != nil {
		log.Println("Site is invalid")
		return err
	}

	if response.Scheme == "" {
		log.Println("Site is invalid")
		return errors.New("Site is invalid")
	}

	if response.Host == "" {
		log.Println("Site is invalid")
		return errors.New("Site is invalid")
	}

	return nil
}
