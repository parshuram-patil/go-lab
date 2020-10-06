package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/servicediscovery"
)

type RegRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	FullName  string `json:"fullName"`
	Title     string `json:"title"`
	Company   string `json:"company"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
}

type GetUserResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	FullName  string `json:"fullName"`
	Title     string `json:"title"`
	Company   string `json:"company"`
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
	case "POST":
		status := r.URL.Query()["status"]
		if len(status) > 0 {
			setHealthStatus(status[0])
		}
	}

	heathStatus := getHealthStatus()
	if heathStatus.Status != CONST_STATUS_UP {
		w.WriteHeader(500)
	}
	json.NewEncoder(w).Encode(getHealthStatus())
}

func getEvn(varName, defaultvalue string) string {
	varValue := os.Getenv(varName)
	if varValue != "" {
		return varValue
	} else {
		return defaultvalue
	}
}

func HandleCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	if r.Method == "OPTIONS" {
		return
	}
}

type ServiceDiscoryResponse struct {
	InstanceIp   string `json:"instanceIp"`
	InstancePort string `json:"instancePort"`
}

func discoverSerive(serviceName, namespaceName string) (*ServiceDiscoryResponse, *ErrorResponse) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
	})
	if err != nil {
		fmt.Println("Error Creating AWS session")
		errMsg := err.Error()
		fmt.Println(errMsg)
		errResponse := ErrorResponse{
			Error: string(errMsg),
		}

		return nil, &errResponse
	}

	svc := servicediscovery.New(sess)
	input := &servicediscovery.DiscoverInstancesInput{
		HealthStatus:  aws.String("HEALTHY"),
		MaxResults:    aws.Int64(100),
		NamespaceName: aws.String(namespaceName),
		ServiceName:   aws.String(serviceName),
	}

	result, err := svc.DiscoverInstances(input)
	if err != nil {
		fmt.Println("Error in Discover instance for " + serviceName + " in " + namespaceName)
		errMsg := err.Error()
		fmt.Println(errMsg)
		errResponse := ErrorResponse{
			Error: string(errMsg),
		}

		return nil, &errResponse
	}

	rand.Seed(time.Now().UnixNano())
	noOfDiscoverdInstances := len(result.Instances)
	attributes := result.Instances[rand.Intn(noOfDiscoverdInstances)].Attributes
	return &ServiceDiscoryResponse{
		InstanceIp:   *attributes["AWS_INSTANCE_IPV4"],
		InstancePort: *attributes["AWS_INSTANCE_PORT"],
	}, nil
}
