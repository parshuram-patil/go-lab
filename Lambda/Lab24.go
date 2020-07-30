package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(hello)
}

type MyEvent struct {
	Name string `json:"uname"`
	Age  int    `json:"uage"`
}

type MyResponse struct {
	Message string `json:"answer"`
}

func hello(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}
