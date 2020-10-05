package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	HandleCORS(w, r)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body")
		errMsg := err.Error()
		fmt.Println(errMsg)
		w.WriteHeader(500)
		errResponse := ErrorResponse{
			Error: string(errMsg),
		}
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	var regRequest RegRequest
	json.Unmarshal(reqBody, &regRequest)
	regRequest.FullName = string(regRequest.FirstName + " " + regRequest.LastName)
	regRequest.Email = string(regRequest.FirstName + "." + regRequest.LastName + "@gmail.com")

	reqBody, err = json.Marshal(regRequest)
	if err != nil {
		fmt.Println("Error preparing DAO data")
		errMsg := err.Error()
		fmt.Println(errMsg)
		w.WriteHeader(500)
		errResponse := ErrorResponse{
			Error: string(errMsg),
		}
		json.NewEncoder(w).Encode(errResponse)
		return
	}

	var DB_SERVICE_HOST = getEvn("DB_SERVICE_HOST", "localhost")
	var DB_SERVICE_PORT = getEvn("DB_SERVICE_PORT", "8092")
	resp, err := http.Post("http://"+DB_SERVICE_HOST+":"+DB_SERVICE_PORT+"/registration", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error calling registration API")
		errMsg := err.Error()
		fmt.Println(errMsg)
		w.WriteHeader(500)
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
		fmt.Println(errMsg)
		w.WriteHeader(500)
		errResponse := ErrorResponse{
			Error: string(errMsg),
		}
		json.NewEncoder(w).Encode(errResponse)
	} else {
		json.NewEncoder(w).Encode(regRequest)
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	HandleCORS(w, r)
	email := r.URL.Query()["email"]
	if len(email) > 0 {
		var DB_SERVICE_HOST string
		var DB_SERVICE_PORT string
		local := r.URL.Query()["local"]
		localValue, _ := strconv.ParseBool(local[0])
		if len(local) > 0 && localValue {
			DB_SERVICE_HOST = "localhost"
			DB_SERVICE_PORT = "8092"
		} else {
			serviceName := getEvn("DB_SERVICE_NAME", "psp-db-api-service")
			namespaceName := getEvn("DB_SERVICE_NAMESPACE", "local")
			dsResult, dsErr := discoverSerive(serviceName, namespaceName)
			if dsErr != nil {
				fmt.Println("Error calling Service Discovery Util")
				w.WriteHeader(500)
				json.NewEncoder(w).Encode(dsErr)
				return
			}
			DB_SERVICE_HOST = dsResult.InstanceIp
			DB_SERVICE_PORT = dsResult.InstancePort
		}

		resp, err := http.Get("http://" + DB_SERVICE_HOST + ":" + DB_SERVICE_PORT + "/user?email=" + email[0])

		if err != nil {
			fmt.Println("Error calling user API")
			errMsg := err.Error()
			fmt.Println(errMsg)
			w.WriteHeader(500)
			errResponse := ErrorResponse{
				Error: string(errMsg),
			}
			json.NewEncoder(w).Encode(errResponse)
			return
		}

		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)

		if resp.StatusCode != 200 {
			fmt.Println("Error from DB Service")
			w.WriteHeader(500)
			fmt.Fprint(w, string(respBody))
			return
		}

		if err != nil {
			fmt.Println("Error reading GetUSer response")
			errMsg := err.Error()
			fmt.Println(errMsg)
			w.WriteHeader(500)
			errResponse := ErrorResponse{
				Error: string(errMsg),
			}
			json.NewEncoder(w).Encode(errResponse)
			return
		}

		var user GetUserResponse
		json.Unmarshal(respBody, &user)
		json.NewEncoder(w).Encode(user)

	} else {
		w.WriteHeader(500)
		errResponse := ErrorResponse{
			Error: "No User Email found in URL",
		}
		json.NewEncoder(w).Encode(errResponse)
	}

}

func ServiceDiscoveryHandler(w http.ResponseWriter, r *http.Request) {
	HandleCORS(w, r)
	var SERVICE_NAME = getEvn("SERVICE_NAME", "psp-db-api-service")
	var NAMESPACE_NAME = getEvn("NAMESPACE_NAME", "local")
	result, err := discoverSerive(SERVICE_NAME, NAMESPACE_NAME)
	if err != nil {
		fmt.Println("Error calling Service Discovery Util")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func registraionService() {
	http.HandleFunc("/registration", RegistrationHandler)
	http.HandleFunc("/health", HeathHandler)
	http.HandleFunc("/user", UserHandler)
	http.HandleFunc("/service", ServiceDiscoveryHandler)
	var HOST_PORT = getEvn("HOST_PORT", "8091")
	fmt.Printf("sever starting on " + HOST_PORT + "\n")
	log.Fatal(http.ListenAndServe(":"+HOST_PORT, nil))
}

func main() {
	registraionService()
}
