package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// files, err := ioutil.ReadDir(".")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, f := range files {
	// 	fmt.Printf("%v - %v - %v \n", f.Name(), f.Size(), f.IsDir())
	// }

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1"),
	})

	if err != nil {
		fmt.Println("In err", err)
		panic("Problem in getting connection to AWS")
	} else {
		svc := s3.New(sess)
		input := &s3.ListObjectsInput{
			Bucket:  aws.String("psp-cross-acc-test"),
			MaxKeys: aws.Int64(2),
		}

		// result, err := svc.ListObjects(input)
		// if err != nil {
		// 	if aerr, ok := err.(awserr.Error); ok {
		// 		switch aerr.Code() {
		// 		case s3.ErrCodeNoSuchBucket:
		// 			fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
		// 		default:
		// 			fmt.Println(aerr.Error())
		// 		}
		// 	} else {
		// 		// Print the error, cast err to awserr.Error to get the Code and
		// 		// Message from an error.
		// 		fmt.Println(err.Error())
		// 	}
		// 	return
		// }

		// fmt.Println(result)

		pageCnt := 0
		err := svc.ListObjectsPages(input,
			func(page *s3.ListObjectsOutput, lastPage bool) bool {
				pageCnt++
				fmt.Println(page)
				// reading only first Page
				// return pageCnt < 1
				return true
			})
		if err != nil {
			fmt.Println("Error Paginating bueckt")
		} else {
			fmt.Printf("\n\n Done reading %d pages\n", pageCnt)
		}
	}
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
