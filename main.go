package main

import (
	"project/TodoServer/app"
	"project/TodoServer/model"
)

func main() {
	model.Init()
	app.Run()
}
