package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func RegistrationDaoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var regRequest RegRequest
	json.Unmarshal(reqBody, &regRequest)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1"),
	})
	if err != nil {
		fmt.Println("Error Creating AWS session")
		fmt.Println(err.Error())
	}

	regData, err := dynamodbattribute.MarshalMap(regRequest)
	input := &dynamodb.PutItemInput{
		Item:      regData,
		TableName: aws.String("registration"),
	}

	svc := dynamodb.New(sess)
	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Error inserting record")
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

func registraionDaoService() {
	http.HandleFunc("/registration", RegistrationDaoHandler)
	http.HandleFunc("/health", HeathHandler)
	HOST_PORT := getEvn("HOST_PORT", "8092")
	fmt.Printf("sever starting on " + HOST_PORT + "\n")
	log.Fatal(http.ListenAndServe(":"+HOST_PORT, nil))
}

func main() {
	registraionDaoService()
}
