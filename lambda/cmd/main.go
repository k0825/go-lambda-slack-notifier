package main

import (
	"context"
	"go-lambda-slack-notifier/slack"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	webhookUrl := os.Getenv("SLACK_WEBHOOK_URL")
	if slack.PostMessage(webhookUrl, "Hello, "+name.Name) != nil {
		return "Error", nil
	} else {
		return "Success", nil
	}
}

func main() {
	lambda.Start(HandleRequest)
}
