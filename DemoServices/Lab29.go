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
	HandleCORS(w, r)
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

func UserHandler(w http.ResponseWriter, r *http.Request) {
	HandleCORS(w, r)
	email := r.URL.Query()["email"]
	if len(email) > 0 {
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("eu-west-1"),
		})
		if err != nil {
			fmt.Println("Error Creating AWS session")
			fmt.Println(err.Error())
		}

		svc := dynamodb.New(sess)
		result, err := svc.GetItem(&dynamodb.GetItemInput{
			TableName: aws.String("registration"),
			Key: map[string]*dynamodb.AttributeValue{
				"email": {
					S: aws.String(email[0]),
				},
			},
		})
		if err != nil {
			errMsg := err.Error()
			fmt.Println(errMsg)
			w.WriteHeader(500)
			errResponse := ErrorResponse{
				Error: string(errMsg),
			}
			json.NewEncoder(w).Encode(errResponse)
			return
		}
		if result.Item == nil {
			msg := "Could not find '" + email[0] + "'"
			errResponse := ErrorResponse{
				Error: string(msg),
			}
			json.NewEncoder(w).Encode(errResponse)
		}

		var userRecord GetUserResponse
		err = dynamodbattribute.UnmarshalMap(result.Item, &userRecord)
		if err != nil {
			errMsg := err.Error()
			fmt.Println(errMsg)
			w.WriteHeader(500)
			errResponse := ErrorResponse{
				Error: string(errMsg),
			}
			json.NewEncoder(w).Encode(errResponse)
		} else {
			json.NewEncoder(w).Encode(userRecord)
		}
	} else {
		w.WriteHeader(500)
		errResponse := ErrorResponse{
			Error: "No User Email found in URL",
		}
		json.NewEncoder(w).Encode(errResponse)
	}
}

func registraionDaoService() {
	http.HandleFunc("/registration", RegistrationDaoHandler)
	http.HandleFunc("/health", HeathHandler)
	http.HandleFunc("/user", UserHandler)
	HOST_PORT := getEvn("HOST_PORT", "8092")
	fmt.Printf("sever starting on " + HOST_PORT + "\n")
	log.Fatal(http.ListenAndServe(":"+HOST_PORT, nil))
}

func main() {
	registraionDaoService()
}
