package model

import "github.com/google/uuid"

type (
	// Represents the all Hotels
	Hotel struct {
		HotelID     uuid.UUID `json:"uuid"`
		Name        string    `json:"name"`
		Address     Address   `json:"address"`
		Rating      int       `json:"rating"`
		Description string    `json:"description"`
		Daily       string    `json:"dailyPrice"`
		Rooms       int       `json:"rooms"`
	}
	// Represents the address of a single Hotel
	Address struct {
		Street     string `json:"street"`
		Number     int    `json:"number"`
		City       string `json:"city"`
		State      string `json:"state"`
		PostalCode int    `json:"cep"`
		Country    string `json:"country"`
	}
)
