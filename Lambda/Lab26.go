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
		ContainerOverrides: &batch.ContainerOverrides{
			Environment: []*batch.KeyValuePair{
				&batch.KeyValuePair{
					Name:  aws.String("IS_RUNNING_ON_CLOUD"),
					Value: aws.String("True"),
				},
				&batch.KeyValuePair{
					Name:  aws.String("INPUT_S3_BUCKET_NAME"),
					Value: aws.String(os.Getenv("INPUT_S3_BUCKET_NAME")),
				},
				&batch.KeyValuePair{
					Name:  aws.String("OUTPUT_S3_BUCKET_NAME"),
					Value: aws.String(os.Getenv("OUTPUT_S3_BUCKET_NAME")),
				},
			},
		},
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

// func test() {
// 	input := &batch.SubmitJobInput{
// 		JobName:       aws.String("psp-job-lambda"),
// 		JobDefinition: aws.String(os.Getenv("JOB_DEFINATION")),
// 		JobQueue:      aws.String(os.Getenv("JOB_QUEUE")),
// 		ContainerOverrides: &batch.ContainerOverrides{
// 			Environment: []*batch.KeyValuePair{
// 				&batch.KeyValuePair{
// 					Name:  aws.String("IS_RUNNING_ON_CLOUD"),
// 					Value: aws.String("True"),
// 				},
// 				&batch.KeyValuePair{
// 					Name:  aws.String("INPUT_S3_BUCKET_NAME"),
// 					Value: aws.String("psp-data-bucket"),
// 				},
// 				&batch.KeyValuePair{
// 					Name:  aws.String("OUTPUT_S3_BUCKET_NAME"),
// 					Value: aws.String("psp-result-bucket"),
// 				},
// 			},
// 		},
// 	}

// 	fmt.Printf("%+v\n", input)
// }

func main() {
	lambda.Start(triggerBatchJob)
	//test()
}
