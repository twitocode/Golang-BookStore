package main

import (
	"fmt"

	"github.com/twitocode/go-web/config"
	"github.com/twitocode/go-web/internal/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config, err := config.LoadViperConfig(".")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.DbHost, config.DbUsername, config.DbPassword, config.DbName, config.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Databse Error:%v", err))
	}

	port := fmt.Sprintf("%s%s", ":", config.Port)
	s := server.NewServer(port, db)

	fmt.Printf("Listening on port %s", port)
	if err := s.Run(); err != nil {
		panic(fmt.Sprintf("Could not start the server: %v", err))
	}
}
