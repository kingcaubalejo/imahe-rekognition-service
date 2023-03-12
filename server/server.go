package server

import (
	"log"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"

	"image-rekognition-service/routers"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	c := cors.New(cors.Options{
		AllowCredentials: true,
		OptionsPassthrough: false,
		AllowedHeaders: []string{"Origin", "Authorization","Access-Control-Allow-Origin", "Content-Type"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "PUT", "OPTIONS"},
		AllowedOrigins: []string{"http://localhost:4200/", "*"},
	})

	a.Router = routers.LoadRouter()

	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(a.Router)
}

func (a *App) Run(port string) {
	fmt.Print("http://localhost" + port + "\n")
	log.Fatal(http.ListenAndServe(port, a.Router))
}