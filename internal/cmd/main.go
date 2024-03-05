package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/seew0/healed/internal/cmd/server"
	"github.com/seew0/healed/internal/router"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURI := os.Getenv("DB_URI")
	if dbURI == "" {
		log.Fatal("DB_URI must be set")
	}

	engine := gin.Default()
	router := router.NewRouter()

	serv := server.NewServer(port,engine,router)
	serv.Start()
}
