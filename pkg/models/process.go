package models

import (
	"time"

	"github.com/turancan-p/go-portfolio-tracking/pkg/database/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type (
	Process struct {
		CreatedAt time.Time
		Stock     string        `json:"stock"`
		Type      ProcessType   `json:"type"`
		Status    ProcessStatus `json:"status"`
		Count     int16         `json:"count"`
		Price     float32       `json:"price"`
		Fee       float32       `json:"fee"`
		ID        uint          `gorm:"primary key;autoIncrement" json:"id"`
	}

	ProcessType   string
	ProcessStatus string
)

const (
	Buying  ProcessType = "Buying"
	Selling ProcessType = "Selling"
)

const (
	Waiting   ProcessStatus = "Waiting"
	Successed ProcessStatus = "Successed"
	Cancelled ProcessStatus = "Cancelled"
)

func init() {
	postgres.Connect()
	db = postgres.GetDB()
	db.AutoMigrate(&Process{})
}

func Create(process Process) (*Process, error) {
	result := db.Create(&process)
	if result.Error != nil {
		return nil, result.Error
	}
	return &process, nil
}

func Delete(Id int) (*Process, error) {
	var process Process
	result := db.Where("ID=?", Id).Delete(process)
	if result.Error != nil {
		return nil, result.Error
	}
	return &process, nil
}

func GetAll() (*[]Process, error) {
	var Process []Process
	result := db.Find(&Process)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Process, nil
}

func GetById(Id int) (*Process, error) {
	var process Process
	result := db.Where("ID=?", Id).Find(&process)
	if result.Error != nil {
		return nil, result.Error
	}
	return &process, nil
}

func Update(Id int, updatedProcess Process) (*Process, error) {
	var process Process

	result := db.First(&process, Id)
	if result.Error != nil {
		return nil, result.Error
	}

	if updatedProcess.Stock != "" {
		process.Stock = updatedProcess.Stock
	}
	if updatedProcess.Type != "" {
		process.Type = updatedProcess.Type
	}
	if updatedProcess.Status != "" {
		process.Status = updatedProcess.Status
	}
	if updatedProcess.Count != 0 {
		process.Count = updatedProcess.Count
	}
	if updatedProcess.Price != 0 {
		process.Price = updatedProcess.Price
	}
	if updatedProcess.Fee != 0 {
		process.Fee = updatedProcess.Fee
	}

	result = db.Save(&process)
	if result.Error != nil {
		return nil, result.Error
	}

	return &process, nil
}
