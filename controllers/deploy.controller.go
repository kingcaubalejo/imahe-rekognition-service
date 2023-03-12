package controllers

import (
	"fmt"
	"net/http"
	"log"
	_"os"
	"io/ioutil"
	"bytes"

	"image-rekognition-service/utils"
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	AWS_S3_REGION = "ap-southeast-1"
	AWS_S3_BUCKET = "img-rekognition-bk"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	session, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION)})
	if err != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": err,
		}, 500, w)
		return
	}

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": err,
			"mod": "form_file",
		}, 500, w)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)
	
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }

	// S3 Upload to bucket
	upload, err := s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key: aws.String("img-rekognition-scan.jpg"),
		ACL: aws.String("private"),
		Body: bytes.NewReader(fileBytes),
		ContentLength: aws.Int64(handler.Size),
		ContentType: aws.String(http.DetectContentType(fileBytes)),
		ContentDisposition: aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	if err != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": err,
			"mod": "s3_file",
		}, 500, w)
		return
	}

	log.Println(upload)

	utils.Response(map[string]interface{}{
		"statusCode": 200,
		"devMessage": ": File upload successful!",
	}, 200, w)
}