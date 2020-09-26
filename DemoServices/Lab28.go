package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body")
		fmt.Println(err.Error())
	}

	var regRequest RegRequest
	json.Unmarshal(reqBody, &regRequest)
	regRequest.FullName = string(regRequest.FirstName + " " + regRequest.LastName)
	regRequest.Email = string(regRequest.FirstName + "." + regRequest.LastName + "@gmail.com")

	reqBody, err = json.Marshal(regRequest)
	if err != nil {
		fmt.Println("Error preparing DAO data")
		fmt.Println(err.Error())
	}

	var DB_SERVICE_HOST = getEvn("DB_SERVICE_HOST", "localhost")
	var DB_SERVICE_PORT = getEvn("DB_SERVICE_PORT", "8092")
	resp, err := http.Post("http://"+DB_SERVICE_HOST+":"+DB_SERVICE_PORT+"/registration", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error calling DAO service")
		errMsg := err.Error()
		fmt.Println(errMsg)
		errResponse := ErrorResponse{
			Error: string(errMsg),
		}
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error DAO response")
		errMsg := err.Error()
		fmt.Println(err.Error())
		w.WriteHeader(500)
		errResponse := ErrorResponse{
			Error: string(errMsg),
		}
		json.NewEncoder(w).Encode(errResponse)
	} else {
		json.NewEncoder(w).Encode(regRequest)
	}
}

func registraionService() {
	http.HandleFunc("/registration", RegistrationHandler)
	http.HandleFunc("/health", HeathHandler)
	var HOST_PORT = getEvn("HOST_PORT", "8091")
	fmt.Printf("sever starting on " + HOST_PORT + "\n")
	log.Fatal(http.ListenAndServe(":"+HOST_PORT, nil))
}

func main() {
	registraionService()
}
