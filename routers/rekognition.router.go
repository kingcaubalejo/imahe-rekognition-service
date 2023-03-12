package routers

import (
	"github.com/gorilla/mux"
	_"github.com/urfave/negroni"

	"image-rekognition-service/controllers"
)

func RekognitionRoute(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/rekognize", controllers.Rekog).Methods("POST")
	return router
}