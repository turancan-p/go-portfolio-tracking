package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/turancan-p/go-portfolio-tracking/pkg/controllers"
)

func SetRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/process", controllers.CreateProcess)
	api.Get("/process", controllers.GetAllProcess)
	api.Get("/process/:id", controllers.GetProcessById)
	api.Put("/process/:id", controllers.UpdateProcess)
	api.Delete("/process/:id", controllers.DeleteProcess)
	log.Println("Router set successfully")
}
