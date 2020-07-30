package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"

	"encoding/json"
	"fmt"
	"os"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")})

	if err != nil {
		fmt.Println("In err", err)
		panic("Problem in getting connection to AWS")
	} else {
		// fmt.Println("Session Created Successfully")
		// fmt.Println(sess)
		client := lambda.New(sess)

		// type MyRequest struct {
		// 	Name string `json:"uname"`
		// 	Age  int    `json:"uage"`
		// }

		// request := MyRequest{
		// 	Name: "Check",
		// 	Age:  3333,
		// }

		payload, err := json.Marshal(map[string]interface{}{
			"uname": "Parshuram",
			"uage":  28,
		})
		result, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String("pspGoFuntions"), Payload: payload})
		if err != nil {
			fmt.Println("Error calling pspGoFuntions : \n", err)
		} else {
			var resp map[string]string
			err = json.Unmarshal(result.Payload, &resp)
			if err != nil {
				fmt.Println("Error unmarshalling MyGetItemsFunction response")
				os.Exit(0)
			} else {
				fmt.Println("Funtion Response: ", resp["answer"])
			}

		}

	}

}
