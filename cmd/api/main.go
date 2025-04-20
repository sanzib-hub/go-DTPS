package main

import (
	"log"
	"taskprocessor/internal/database/postgres"
	"taskprocessor/cmd/server"

	"github.com/joho/godotenv"
)



func main(){
	err1 := godotenv.Load(".env")
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}
	err := postgres.InitPostgres()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	err = tcpServer.StartTCPServer()
	if err != nil {
		log.Fatalf("TCP Server error: %v", err)
	}
}
