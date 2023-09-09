package helper

import "github.com/turancan-p/go-portfolio-tracking/pkg/models"

func CreateResponseProcess(process models.Process) models.Process {
	return models.Process{ID: process.ID, Stock: process.Stock,
		Type: models.ProcessType(process.Type), Status: models.ProcessStatus(process.Status),
		Count: process.Count, Price: process.Price, Fee: process.Fee}
}
