package controllers

import (
	"image-rekognition-service/utils"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	utils.Response(map[string]interface{}{
		"statusCode": 200,
		"devMessage": "Pong!",
	}, 200, w)
}