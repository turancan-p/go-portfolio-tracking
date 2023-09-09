package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/turancan-p/go-portfolio-tracking/pkg/database/postgres"
	"github.com/turancan-p/go-portfolio-tracking/pkg/helper"
	"github.com/turancan-p/go-portfolio-tracking/pkg/models"
)

func CreateProcess(context *fiber.Ctx) error {
	var process models.Process

	if err := context.BodyParser(&process); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	postgres.Database.Db.Create(&process)
	responseProcess := helper.CreateResponseProcess(process)

	return context.Status(200).JSON(responseProcess)
}

func GetAllProcess(context *fiber.Ctx) error {
	processes := []models.Process{}

	postgres.Database.Db.Find(&processes)

	responseProcesses := []models.Process{}

	for _, process := range processes {
		responseProcess := helper.CreateResponseProcess(process)
		responseProcesses = append(responseProcesses, responseProcess)
	}

	return context.Status(200).JSON(responseProcesses)
}

func findProcess(id int, process *models.Process) error {
	postgres.Database.Db.Find(process, "id = ?", id)
	if process.ID == 0 {
		return errors.New("Process does not exist")
	}
	return nil
}

func GetProcess(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")

	var process models.Process
	if err != nil {
		return context.Status(400).JSON("Please ensure that :id is an integer")
	}
	if err := findProcess(id, &process); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	responseProcess := helper.CreateResponseProcess(process)
	return context.Status(200).JSON(responseProcess)
}

func DeleteProcess(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")

	var process models.Process
	if err != nil {
		return context.Status(400).JSON("Please ensure that :id is an integer")
	}
	if err := findProcess(id, &process); err != nil {
		return context.Status(400).JSON(err.Error())
	}
	if err := postgres.Database.Db.Delete(&process).Error; err != nil {
		return context.Status(404).JSON(err.Error())
	}
	return context.Status(200).SendString("Successfully deleted process")
}

func UpdateProcess(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")

	var process models.Process
	if err != nil {
		return context.Status(400).JSON("Please ensure that :id is an integer")
	}
	if err := findProcess(id, &process); err != nil {
		return context.Status(400).JSON(err.Error())
	}
	type UpdateProcess struct {
		Stock  string               `json:"stock"`
		Type   models.ProcessType   `json:"type"`
		Status models.ProcessStatus `json:"status"`
		Count  int16                `json:"count"`
		Price  float32              `json:"price"`
		Fee    float32              `json:"fee"`
	}
	var updateData UpdateProcess

	if err := context.BodyParser(&updateData); err != nil {
		return context.Status(500).JSON(err.Error())
	}
	if updateData.Stock != "" {
		process.Stock = updateData.Stock
	}
	if updateData.Type != "" {
		process.Type = updateData.Type
	}
	if updateData.Status != "" {

		process.Status = updateData.Status
	}
	if updateData.Count != 0 {
		process.Count = updateData.Count
	}
	if updateData.Price != 0 {
		process.Price = updateData.Price
	}
	if updateData.Fee != 0 {
		process.Fee = updateData.Fee
	}

	postgres.Database.Db.Save(&process)

	responseProcess := helper.CreateResponseProcess(process)

	return context.Status(200).JSON(responseProcess)
}
