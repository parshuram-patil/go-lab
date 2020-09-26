package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"sync"
)

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

var (
	CONST_STATUS_UP   string = "UP"
	CONST_STATUS_DOWN string = "DOWN"
)

var currentHealthStatus string = CONST_STATUS_UP
var healthMutex = &sync.Mutex{}

type HealthStatus struct {
	Status string `json:"status"`
}

func getHealthStatus() HealthStatus {

	return HealthStatus{
		Status: currentHealthStatus,
	}
}

func setHealthStatus(status string) HealthStatus {
	healthMutex.Lock()
	status = strings.ToUpper(status)
	if status == CONST_STATUS_UP || status == CONST_STATUS_DOWN {
		currentHealthStatus = status
	}
	healthMutex.Unlock()

	return getHealthStatus()
}

func HeathHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {

	case "GET":
		json.NewEncoder(w).Encode(getHealthStatus())

	case "POST":
		status := r.URL.Query()["status"]
		if len(status) > 0 {
			setHealthStatus(status[0])
			os.Exit(2)
		}
		json.NewEncoder(w).Encode(getHealthStatus())
	}
}

func getEvn(varName, defaultvalue string) string {
	varValue := os.Getenv(varName)
	if varValue != "" {
		return varValue
	} else {
		return defaultvalue
	}
}
