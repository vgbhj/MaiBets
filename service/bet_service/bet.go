package bet_service

import "github.com/vgbhj/MaiBets/models"

func AddBet(bet models.Bet) error {
	mat := map[string]interface{}{
		"client_id":   bet.ClientID,
		"event_id":    bet.EventID,
		"bet_type_id": bet.BetTypeID,
		"odd_id":      bet.OddID,
		"bet_amount":  bet.BetAmount,
		"status":      bet.Status,
		"bet_date":    bet.BetDate,
	}
	return models.AddBet(mat)
}
