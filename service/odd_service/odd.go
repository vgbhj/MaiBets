package odd_service

import "github.com/vgbhj/MaiBets/models"

// AddOdd добавляет новую ставку
func AddOdd(odd models.Odd) error {
	mat := map[string]interface{}{
		"id":         odd.ID,
		"odd_value":  odd.OddValue,
		"event_id":   odd.EventID,
		"updated_at": odd.UpdatedAt,
	}
	return models.AddOdd(mat)
}

// GetOdd получает ставку по ее ID
func GetOdd(id int) (*models.Odd, error) {
	material, err := models.GetOdd(id)
	if err != nil {
		return nil, err
	}
	return material, nil
}
