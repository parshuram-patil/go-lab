package main

import "os"

type RegRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	FullName  string `json:"fullName"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
}

type RegResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	FullName  string `json:"fullName"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func getEvn(varName, defaultvalue string) string {
	varValue := os.Getenv(varName)
	if varValue != "" {
		return varValue
	} else {
		return defaultvalue
	}
}
