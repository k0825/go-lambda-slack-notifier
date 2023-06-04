package utils

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestGetBottomDirectory(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"パスが正しい時", args{"hoge/fuga/PC/piyo.jpeg"}, "PC", false},
		{"パスが不正な時", args{"/"}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBottomDirectory(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBottomDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				fmt.Println(got, tt.want)
				t.Errorf("GetBottomDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateMessageFromS3EventRecord(t *testing.T) {
	type args struct {
		record events.S3EventRecord
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"保存先がPCディレクトリの配下の時", args{events.S3EventRecord{
			S3: events.S3Entity{
				Bucket: events.S3Bucket{
					Name: "hoge",
				},
				Object: events.S3Object{
					Key: "hoge/fuga/PC/piyo.jpeg",
				},
			},
		}}, "New file uploaded(PC): https://hoge.s3.amazonaws.com/hoge/fuga/PC/piyo.jpeg", false},
		{"保存先がSMPディレクトリの配下の時", args{events.S3EventRecord{
			S3: events.S3Entity{
				Bucket: events.S3Bucket{
					Name: "hoge",
				},
				Object: events.S3Object{
					Key: "hoge/fuga/SMP/piyo.jpeg",
				},
			},
		}}, "New file uploaded(SMP): https://hoge.s3.amazonaws.com/hoge/fuga/SMP/piyo.jpeg", false},
		{"保存先がPCかSMPディレクトリ以外の時", args{events.S3EventRecord{
			S3: events.S3Entity{
				Bucket: events.S3Bucket{
					Name: "hoge",
				},
				Object: events.S3Object{
					Key: "hoge/fuga/piyo.jpeg",
				},
			},
		}}, "", true},
		{"パスが不正な時", args{events.S3EventRecord{
			S3: events.S3Entity{
				Bucket: events.S3Bucket{
					Name: "hoge",
				},
				Object: events.S3Object{
					Key: "",
				},
			},
		}}, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateMessageFromS3EventRecord(tt.args.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBottomDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				fmt.Println(got, tt.want)
				t.Errorf("GetBottomDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}
