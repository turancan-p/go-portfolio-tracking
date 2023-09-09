package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/turancan-p/go-portfolio-tracking/pkg/database/postgres"
	"github.com/turancan-p/go-portfolio-tracking/pkg/router"
)

func main() {
	app := fiber.New()
	postgres.NewConnection()

	router.SetRoutes(app)
	log.Println(fmt.Sprintf("API STARTED ON PORT:%s", os.Getenv("API_PORT")))
	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("API_PORT"))))
}
