package server

import (
	"fmt"
	"log"

	"github.com/dreezy305/library-core-service/internal/config"
	"github.com/dreezy305/library-core-service/internal/database"
	"github.com/dreezy305/library-core-service/internal/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func Start() {
	cfg := config.Load()
	database, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	app := fiber.New()
	app.Use(cors.New())

	api := app.Group(cfg.Path)
	fmt.Println(cfg.Path)
	routes.HealthCheckRoute(app)
	routes.AuthRoutes(api, database)
	routes.UserRoutes(api, database)
	routes.AuthorRoutes(api, database)
	routes.BookRoutes(api, database)
	routes.BookCategoryRoutes(api, database)
	// routes.LoanRoutes(api)

	// Implementation of server start logic goes here
	fmt.Println("Server started")
	// Start the server on port 3000
	log.Fatal(app.Listen(":" + cfg.Port))
}
