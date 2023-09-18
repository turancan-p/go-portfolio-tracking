package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/turancan-p/go-portfolio-tracking/pkg/models"
)

func CreateProcess(context *fiber.Ctx) error {
	var process models.Process

	if err := context.BodyParser(&process); err != nil {
		return context.Status(http.StatusBadRequest).JSON(err.Error())
	}

	res, err := models.Create(process)
	if err != nil {
		return context.Status(http.StatusExpectationFailed).JSON(err)
	}
	return context.Status(http.StatusOK).JSON(res)
}

func GetAllProcess(context *fiber.Ctx) error {
	res, err := models.GetAll()
	if err != nil {
		return context.Status(http.StatusNotFound).JSON(err.Error())
	}
	return context.Status(http.StatusOK).JSON(res)
}

func GetProcessById(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	res, err := models.GetById(id)
	if err != nil {
		return context.Status(http.StatusNotFound).JSON(err.Error())
	}
	return context.Status(http.StatusOK).JSON(res)
}

func UpdateProcess(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	var updateData models.Process

	if err := context.BodyParser(&updateData); err != nil {
		return context.Status(500).JSON(err.Error())
	}

	res, err := models.Update(id, updateData)

	if err != nil {
		return context.Status(http.StatusNotFound).JSON(err.Error())
	}
	return context.Status(http.StatusOK).JSON(res)
}

func DeleteProcess(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	res, err := models.GetById(id)
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return context.Status(http.StatusOK).JSON(res)
}
