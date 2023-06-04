package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Slack struct {
	Text string `json:"text"`
}

func sendSlackMessage(webhookUrl string, text string) error {
	slack := Slack{
		Text: text,
	}

	params, _ := json.Marshal(slack)
	res, err := http.PostForm(webhookUrl, url.Values{"payload": {string(params)}})
	defer res.Body.Close()

	return err
}

func getBottomDirectory(path string) (string, error) {
	dir, _ := filepath.Split(path)
	dir = strings.TrimSuffix(dir, "/")
	dirs := strings.Split(dir, "/")

	if len(dirs) > 0 {
		return dirs[len(dirs)-1], nil
	}
	return "", fmt.Errorf("Invalid path")
}

func createMessageFromS3EventRecord(record events.S3EventRecord) (string, error) {
	s3rec := record.S3
	bucketName := s3rec.Bucket.Name
	objKey := s3rec.Object.Key
	objUrl := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, objKey)

	if dir, err := getBottomDirectory(objKey); err != nil {
		return "", err
	} else if dir == "PC" || dir == "SMP" {
		return fmt.Sprintf("New file uploaded(%s): %s", dir, objUrl), nil
	} else {
		return "", fmt.Errorf("Should be put under PC or SMP directory")
	}
}

func HandleRequest(ctx context.Context, s3Event events.S3Event) (string, error) {
	webhookUrl := os.Getenv("SLACK_WEBHOOK_URL")
	var message string

	for _, record := range s3Event.Records {
		if msg, err := createMessageFromS3EventRecord(record); err != nil {
			return "Error", err
		} else {
			message += msg + "\n"
		}
	}

	if err := sendSlackMessage(webhookUrl, message); err != nil {
		return "Error", err
	} else {
		return "Success", nil
	}
}

func main() {
	lambda.Start(HandleRequest)
}
