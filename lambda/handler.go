package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// handler is a simple function that takes a string and does a ToUpper.
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	mySession := session.Must(session.NewSession())

	// Create a DynamoDB client with additional configuration
	svc := dynamodb.New(mySession, aws.NewConfig().WithRegion("ap-northeast-1"))

	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Title": {
				S: aws.String("電気料金"),
			},
			"Date": {
				S: aws.String("2021-08-02"),
			},
		},
		TableName: aws.String("money-log-table"),
	}

	result, err := svc.GetItem(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeRequestLimitExceeded:
				fmt.Println(dynamodb.ErrCodeRequestLimitExceeded, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}

	fmt.Println(result)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       result.String(),
	}, nil
}

func main() {
	lambda.Start(handler)
}
