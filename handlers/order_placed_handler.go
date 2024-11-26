package handlers

import (
	"fmt"

	"github.com/miladsssh/goeventflow/eventbus"
	"github.com/miladsssh/goeventflow/events"
)

func OrderPlacedHandler(event eventbus.Event[events.OrderPlacedEvent]) error {
	fmt.Printf("Handling OrderPlacedEvent: OrderID=%d, Amount=%.2f\n", event.Data.OrderID, event.Data.Amount)
	return nil
}
