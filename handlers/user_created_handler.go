package handlers

import (
	"fmt"

	"github.com/miladsssh/goeventflow/eventbus"
	"github.com/miladsssh/goeventflow/events"
)

func UserCreatedHandler(event eventbus.Event[events.UserCreatedEvent]) error {
	fmt.Printf("Handling UserCreatedEvent: UserID=%d, UserName=%s\n", event.Data.UserID, event.Data.UserName)
	return nil
}
