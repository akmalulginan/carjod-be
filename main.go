package main

import (
	"log"
	"os"

	"github.com/akmalulginan/carjod-be/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	postgres "github.com/akmalulginan/carjod-be/storage/postgres"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	db, err := postgres.Init(postgres.Credential{
		Host:     os.Getenv("POSTGRES_HOST"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	})
	if err != nil {
		panic(err)
	}

	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()

	r := gin.Default()
	routes.Setup(r, db)

}
