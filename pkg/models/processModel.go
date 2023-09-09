package models

import (
	"time"

	"gorm.io/gorm"
)

type ProcessType string

type ProcessStatus string

const (
	Buying  ProcessType   = "Buying"
	Selling ProcessType   = "Selling"
	Waiting ProcessStatus = "Waiting"

	Successed ProcessStatus = "Successed"
	Cancelled ProcessStatus = "Cancelled"
)

type Process struct {
	CreatedAt time.Time
	Stock     string        `json:"stock"`
	Type      ProcessType   `json:"type"`
	Status    ProcessStatus `json:"status"`
	Count     int16         `json:"count"`
	Price     float32       `json:"price"`
	Fee       float32       `json:"fee"`
	ID        uint          `gorm:"primary key;autoIncrement" json:"id"`
}

// for process type
func (p *Process) ChangeType(processType ProcessType) {
	p.Type = processType
}

// for process status
func (p *Process) ChangeStatus(processStatus ProcessStatus) {
	p.Status = processStatus
}

// for db migrations
func DBMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Process{})
	return err
}
