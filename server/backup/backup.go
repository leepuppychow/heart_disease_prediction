package backup

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws/awsutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	S3_REGION = "us-west-2"
	S3_BUCKET = "heart-disease-prediction"
)

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
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccess := os.Getenv("AWS_SECRET_ACCESS_KEY")
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
	path := file.Name()

	params := &s3.PutObjectInput{
		Bucket:        aws.String(S3_BUCKET),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}
	return params
}
