package main

import (
	"developer.patronus.io/mobile_backend/app"
)

func main() {

	app := &app.App{}
	app.Initialize()
	app.Run(":80")
}
