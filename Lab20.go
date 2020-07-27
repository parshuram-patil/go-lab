package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-1"),
		Credentials: credentials.NewSharedCredentials("", "sandbox"),
	})

	if err != nil {
		fmt.Println("In err", err)
	} else {
		svc := dynamodb.New(sess)
		listtablesoutput, err := svc.ListTables(&dynamodb.ListTablesInput{})
		if err != nil {
			fmt.Println("Error ", err)
		} else {
			fmt.Println(listtablesoutput)
		}
	}

}
