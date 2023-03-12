package controllers

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/base64"
	"log"
	"encoding/json"
	"strings"

	"image-rekognition-service/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

type RequestBody struct {
    Image string
}

func Rekog(w http.ResponseWriter, r *http.Request) {

	session, err := session.NewSession(&aws.Config{Region: aws.String(AWS_S3_REGION)})
	if err != nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": err,
		}, 500, w)
		return
	}
	fmt.Println(session)

	body, mgs := utils.HttpReq(r)
	if body == nil {
		utils.Response(map[string]interface{}{
			"statusCode": 500,
			"devMessage": mgs,
		}, 500, w)
		return
	}

	
	requestPayload, _ := body["imgData"].(string)
	b64data := requestPayload[strings.IndexByte(requestPayload, ',')+1:]

	decodedImage, err := base64.StdEncoding.DecodeString(b64data)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
	
	input := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image {
			Bytes: decodedImage,
			// S3Object: &rekognition.S3Object {
			// 	Bucket: aws.String("img-rekognition-bk"),
			// 	Name: aws.String("WIN_20230217_13_35_00_Pro.jpg"),
			// },
		},
	}
	svc := rekognition.New(session)
	result, err := svc.DetectLabels(input)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(result)
	fmt.Println("Output: ", output)

	utils.Response(map[string]interface{}{
		"statusCode": 200,
		"devMessage": result,
	}, 200, w)
}

func encode(bin []byte) []byte {
	e64 := base64.StdEncoding

	maxEncLen := e64.EncodedLen(len(bin))
	encBuf := make([]byte, maxEncLen)

	e64.Encode(encBuf, bin)
	return encBuf
}

func format(enc []byte, mime string) string {
	switch mime {
	case "image/gif", "image/jpeg", "image/pjpeg", "image/png", "image/tiff":
		return fmt.Sprintf("data:%s;base64,%s", mime, enc)
	default:
	}

	return fmt.Sprintf("data:image/png;base64,%s", enc)
}

func FromRemote(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error getting url.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	ct := resp.Header.Get("Content-Type")



	image, mime := body, ct
	enc := encode(image)

	out := format(enc, mime)
	return out
}