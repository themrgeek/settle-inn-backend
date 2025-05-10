package middlewares

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadToS3(file *multipart.FileHeader) (string, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	s3client := s3.New(sess)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	buf := bytes.NewBuffer(nil)
	_, err = buf.ReadFrom(src)
	if err != nil {
		return "", err
	}

	key := "uploads/" + file.Filename
	_, err = s3client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buf.Bytes()),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", os.Getenv("S3_BUCKET"), key)
	return url, nil
}
