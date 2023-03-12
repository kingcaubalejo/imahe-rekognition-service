package routers

import (
	"github.com/gorilla/mux"

	"image-rekognition-service/controllers"
	
)

func PingRoute(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/ping", controllers.Ping).Methods("GET")
	return router
}


func LoadRouter() *mux.Router {
	router := mux.NewRouter()
	router = UploadRoute(router)
	router = RekognitionRoute(router)
	router = PingRoute(router)
	
	return router
}