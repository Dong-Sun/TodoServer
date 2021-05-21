package main

import (
	"TodoServer/app"
	"TodoServer/model"
)

func main() {
	model.Init()
	app.Run()
}
