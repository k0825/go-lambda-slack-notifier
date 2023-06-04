package main

import (
	"context"
	"go-lambda-slack-notifier/utils"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, s3Event events.S3Event) (string, error) {
	webhookUrl := os.Getenv("SLACK_WEBHOOK_URL")
	var message string

	for _, record := range s3Event.Records {
		if msg, err := utils.CreateMessageFromS3EventRecord(record); err != nil {
			return "Error", err
		} else {
			message += msg + "\n"
		}
	}

	if err := utils.SendSlackMessage(webhookUrl, message); err != nil {
		return "Error", err
	} else {
		return "Success", nil
	}
}

func main() {
	lambda.Start(HandleRequest)
}
