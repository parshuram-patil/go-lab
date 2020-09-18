package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/batch"
)

type MyResponse struct {
	Message string `json:"answer"`
}

func triggerBatchJob(ctx context.Context, event events.CloudWatchEvent) (MyResponse, error) {

	svc := batch.New(session.New())
	input := &batch.SubmitJobInput{
		JobName:       aws.String("psp-job-lambda"),
		JobDefinition: aws.String(os.Getenv("JOB_DEFINATION")),
		JobQueue:      aws.String(os.Getenv("JOB_QUEUE")),
	}

	result, err := svc.SubmitJob(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case batch.ErrCodeClientException:
				fmt.Println(batch.ErrCodeClientException, aerr.Error())
			case batch.ErrCodeServerException:
				fmt.Println(batch.ErrCodeServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}

	resultJson, _ := json.MarshalIndent(result, "", "  ")
	response := string(resultJson)
	return MyResponse{Message: response}, nil
}

func main() {
	lambda.Start(triggerBatchJob)
}
