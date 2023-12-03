package main

import (
	"github.com/kitsune2411/ginCRUD/config"
	"github.com/kitsune2411/ginCRUD/models"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&models.Post{})
}
