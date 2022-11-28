package model

import "time"

type Company struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type Student struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Dob         time.Time `json:"dob"`
	PhoneNumber string    `json:"phoneNumber"`
	Branch      string    `json:"branch"`
	Company     Company   `json:"company"`
}
