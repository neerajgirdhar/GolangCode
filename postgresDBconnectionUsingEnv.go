package main

import (
	"awesomeProject/initializers"
	"awesomeProject/model"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	todo := model.Todo{Content: "co", Status: true}
	initializers.Db.Create(&todo)
	//initializers.Db.AutoMigrate(&model.Todo{})
}
