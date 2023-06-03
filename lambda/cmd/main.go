package main

import (
	"context"
	"go-lambda-slack-notifier/slack"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, s3Event events.S3Event) (string, error) {
	webhookUrl := os.Getenv("SLACK_WEBHOOK_URL")

	for _, record := range s3Event.Records {
		s3rec := record.S3
		log.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3rec.Bucket.Name, s3rec.Object.Key)
	}

	if slack.PostMessage(webhookUrl, "Hello") != nil {
		return "Error", nil
	} else {
		return "Success", nil
	}
}

func main() {
	lambda.Start(HandleRequest)
}
