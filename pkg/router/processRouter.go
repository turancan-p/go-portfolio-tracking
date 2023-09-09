package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/turancan-p/go-portfolio-tracking/pkg/handlers"
)

func SetRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/process", handlers.CreateProcess)
	api.Get("/process", handlers.GetAllProcess)
	api.Get("/process/:id", handlers.GetProcess)
	api.Put("/process/:id", handlers.UpdateProcess)
	api.Delete("/process/:id", handlers.DeleteProcess)
	log.Println("Router set successfully")
}
