package main

import (
	"github.com/skondrag/go-curd/initializers"
	"github.com/skondrag/go-curd/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
