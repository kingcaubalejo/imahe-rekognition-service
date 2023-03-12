package main

import "image-rekognition-service/server"

func main() {
	app := server.App{}
	app.Initialize()
	app.Run(":8081")
}
