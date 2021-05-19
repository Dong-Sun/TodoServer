package main

import (
	"edu/TodoServer/app"
	"edu/TodoServer/model"
)

func main() {
	model.Init()
	app.Run()
}
