package routers

import (
	"github.com/gorilla/mux"
	_"github.com/urfave/negroni"

	"image-rekognition-service/controllers"
)

func UploadRoute(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/upload", controllers.Upload).Methods("PUT")
	return router
}