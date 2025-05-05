package subscriber

import (
	"encoding/json"
	"log"
	"statistics_service/internal/usecase"
	"time"

	"github.com/nats-io/nats.go"
)

type OrderCreatedEvent struct {
	UserID    string  `json:"user_id"`
	Amount    float64 `json:"amount"`
	Timestamp int64   `json:"timestamp"`
}

func StartNATSSubscriber(uc *usecase.StatisticsUsecase) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	_, err = nc.Subscribe("order.created", func(m *nats.Msg) {
		var event OrderCreatedEvent
		if err := json.Unmarshal(m.Data, &event); err != nil {
			log.Println("Failed to unmarshal order.created event:", err)
			return
		}
		hour := time.Unix(event.Timestamp, 0).Hour()
		if err := uc.UpdateUserStatistics(event.UserID, event.Amount, hour); err != nil {
			log.Println("Failed to update user stats:", err)
		}
	})

	log.Println("Subscribed to NATS topic: order.created")
	select {}
}
