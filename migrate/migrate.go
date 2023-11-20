package main

import (
	"github.com/mxnuchim/golang-crud-api/initializers"
	"github.com/mxnuchim/golang-crud-api/models"
)

func  init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}