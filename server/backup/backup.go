package backup

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var S3_REGION = os.Getenv("S3_REGION")
var S3_BUCKET = os.Getenv("S3_BUCKET")
var accessKey = os.Getenv("AWS_ACCESS_KEY_ID")
var secretAccess = os.Getenv("AWS_SECRET_ACCESS_KEY")

func SaveToS3(filepath string) error {
	creds, _ := getCredentials()
	cfg := aws.NewConfig().WithRegion(S3_REGION).WithCredentials(creds)
	svc := s3.New(session.New(), cfg)
	params := getFileParams(filepath)
	resp, err := svc.PutObject(params)
	if err != nil {
		log.Println("Error uploading file to S3", err)
	} else {
		log.Println("Successfully loaded file to S3", awsutil.StringValue(resp))
	}
	return err
}

func getCredentials() (*credentials.Credentials, error) {
	token := ""
	creds := credentials.NewStaticCredentials(accessKey, secretAccess, token)

	_, err := creds.Get()
	if err != nil {
		log.Println("error getting AWS credentials", err)
	}
	return creds, err
}

func getFileParams(filepath string) *s3.PutObjectInput {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Println("Error opening file for S3 upload", err)
	}
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	path := addTimeStamp(file.Name())

	params := &s3.PutObjectInput{
		Bucket:        aws.String(S3_BUCKET),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}
	return params
}

func addTimeStamp(filename string) string {
	stamp := strconv.Itoa(int(time.Now().Unix())) + ".csv"
	return strings.Replace(filename, ".csv", stamp, 1)
}
