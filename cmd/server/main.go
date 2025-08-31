package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vinihss/bodego-api/config"
	_ "github.com/vinihss/bodego-api/docs"
	"github.com/vinihss/bodego-api/internal/infrastructure/database/models"
	"github.com/vinihss/bodego-api/internal/routes"
	"log"
	"os"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) Run() error {
	config.ConnectDB()
	err := config.DB.AutoMigrate(models.Customer{}, models.Drink{})
	if err != nil {
		return errors.New(fmt.Sprintf("Error migrating database: %v", err))
	}

	// Seed drinks
	if err := models.SeedDrinks(config.DB); err != nil {
		return errors.New(fmt.Sprintf("Error seeding drinks: %v", err))
	}

	// Set up Gin router
	r := gin.Default()
	routes.SetupRoutes(r)

	err = r.Run(s.addr)
	if err != nil {
		return errors.New(fmt.Sprintf("Error starting server: %v", err))
	}
	// Start the server
	return nil
}

// @title Favorite API
// @version 1.0
// @description API para favoritar produtos
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @BasePath /
func main() {

	// load env from os

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := NewServer(":" + port)
	err := server.Run()
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
		os.Exit(1)
	}
}
