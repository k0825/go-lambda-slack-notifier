package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type Slack struct {
	Text string `json:"text"`
}

func SendSlackMessage(webhookUrl string, text string) error {
	slack := Slack{
		Text: text,
	}

	params, _ := json.Marshal(slack)
	res, err := http.PostForm(webhookUrl, url.Values{"payload": {string(params)}})
	defer res.Body.Close()

	return err
}

func GetBottomDirectory(path string) (string, error) {
	dir, _ := filepath.Split(path)
	dir = strings.TrimSuffix(dir, "/")
	dirs := strings.Split(dir, "/")

	if len(dirs) == 1 && dirs[0] == "" {
		return "", fmt.Errorf("Invalid path")
	}

	return dirs[len(dirs)-1], nil
}

func CreateMessageFromS3EventRecord(record events.S3EventRecord) (string, error) {
	s3rec := record.S3
	bucketName := s3rec.Bucket.Name
	objKey := s3rec.Object.Key
	objUrl := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, objKey)

	if dir, err := GetBottomDirectory(objKey); err != nil {
		return "", err
	} else if dir == "PC" || dir == "SMP" {
		return fmt.Sprintf("New file uploaded(%s): %s", dir, objUrl), nil
	} else {
		return "", fmt.Errorf("Should be put under PC or SMP directory")
	}
}
