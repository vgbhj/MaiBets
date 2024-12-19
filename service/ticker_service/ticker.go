package ticker_service

import (
	"fmt"
	"time"

	"github.com/vgbhj/MaiBets/models"
)

// StartTicker запускает периодическую проверку событий
func StartTicker() {
	go func() {
		ticker := time.NewTicker(10 * time.Second) // Проверка каждую минуту
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				err := models.UpdateEventStatus()
				if err != nil {
					// Логирование ошибки
					fmt.Println("Error updating event status:", err)
				}
			}
		}
	}()
}
