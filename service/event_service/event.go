package event_service

import (
	"github.com/vgbhj/MaiBets/models"
)

func AddEvent(event models.Event) error {
	mat := map[string]interface{}{
		"id":     event.ID,
		"name":   event.Name,
		"desc":   event.Desc,
		"date":   event.Date,
		"status": event.Status,
	}
	return models.AddEvent(mat)
}

func GetEvent(id int) (*models.Event, error) {
	material, err := models.GetEvent(id)
	if err != nil {
		return nil, err
	}
	return material, nil
}
