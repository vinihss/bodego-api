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
	err := config.DB.AutoMigrate(models.Customer{}, models.News{})
	if err != nil {
		return errors.New(fmt.Sprintf("Error migrating database: %v", err))
	}

	// Seed default news data if the table is empty
	var count int64
	config.DB.Model(&models.News{}).Count(&count)
	if count == 0 {
		defaultNews := []models.News{
			{
				Name:        "Cerveja Heineken Latão",
				Price:       10.00,
				Description: "Cerveja Heineken em lata",
			},
			{
				Name:        "Dose cachaça",
				Price:       4.00,
				Description: "Dose de cachaça",
			},
		}
		for _, news := range defaultNews {
			config.DB.Create(&news)
		}
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
