package middlewares

import (
	"fmt"

	"github.com/miladsssh/goeventflow/eventbus"
	"github.com/miladsssh/goeventflow/events"
)

func UsernameValidatorMiddleware(event *eventbus.Event[events.UserCreatedEvent]) (bool, error) {
	if len(event.Data.UserName) < 3 {
		fmt.Println("Username too short, skipping event.")
		return false, nil
	}
	return true, nil
}

func AmountValidatorMiddleware(event *eventbus.Event[events.OrderPlacedEvent]) (bool, error) {
	if event.Data.Amount <= 0 {
		fmt.Println("Invalid order amount, skipping event.")
		return false, nil
	}
	return true, nil
}
