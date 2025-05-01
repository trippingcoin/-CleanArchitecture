package natsadapter

import (
	"encoding/json"
	"statistics_service/internal/app"

	"github.com/nats-io/nats.go"
)

type OrderCreatedEvent struct {
	UserID string
	Amount float64
	Hour   int
}

func Subscribe(nc *nats.Conn, app *app.StatisticsApp) {
	nc.Subscribe("order.created", func(m *nats.Msg) {
		var event OrderCreatedEvent
		if err := json.Unmarshal(m.Data, &event); err == nil {
			app.ProcessOrder(event.UserID, event.Amount, event.Hour)
		}
	})
}
