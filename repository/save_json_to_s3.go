package repository

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func SaveJsonToS3(content []byte) {
	fmt.Println("Start process ...")
	fmt.Println("Download JSON in AWS S3")
	seoul, _ := time.LoadLocation("Asia/Seoul")
	timeData := time.Now().In(seoul)
	item := timeData.String() + ".json"
	bucket := "italian-bmt-bucket"
	accessKeyID := os.Getenv("ACCESSKEYID")
	accessKeySecret := os.Getenv("ACCESSKEYSECRET")

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-2"),
		Credentials: credentials.NewStaticCredentials(accessKeyID, accessKeySecret, ""),
	}))
	svc := s3.New(sess)
	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path.Join(strconv.Itoa(timeData.Year()), strconv.Itoa(int(timeData.Month())), strconv.Itoa(timeData.Day()), item)),
		Body:   bytes.NewReader(content),
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("File uploaded successfully")
	fmt.Println("END process ...")
}
