package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/twitocode/go-web/internal/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(fmt.Sprintf("Could not load env files: %v", err))
	}

	port := fmt.Sprintf("%s%s", ":", os.Getenv("PORT"))

	if port == "" {
		panic("Port not found in env")
	}

	var (
		dbHost     = os.Getenv("DB_HOST")
		dbPort     = os.Getenv("DB_PORT")
		dbName     = os.Getenv("DB_NAME")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbUsername = os.Getenv("DB_USERNAME")
	)
  
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUsername, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Databse Error:%v", err))
	}

	s := server.NewServer(port, db)

	fmt.Printf("Listening on port %s", port)
	if err := s.Run(); err != nil {
		panic(fmt.Sprintf("Could not start the server: %v", err))
	}
}
