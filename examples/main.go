package main

import (
	"fmt"

	"github.com/miladsssh/goeventflow/eventbus"
	"github.com/miladsssh/goeventflow/events"
	"github.com/miladsssh/goeventflow/handlers"
	"github.com/miladsssh/goeventflow/middlewares"
)

func main() {
	bus := eventbus.NewEventBus()

	// Register global middlewares
	bus.RegisterGlobalMiddlewares(middlewares.GlobalLoggingMiddleware)

	// Register middlewares for specific events
	eventbus.RegisterMiddleware(bus, middlewares.LoggingMiddleware[events.UserCreatedEvent], middlewares.UsernameValidatorMiddleware)
	eventbus.RegisterMiddleware(bus, middlewares.LoggingMiddleware[events.OrderPlacedEvent], middlewares.AmountValidatorMiddleware)

	// Register handlers
	eventbus.RegisterHandler(bus, handlers.UserCreatedHandler)
	eventbus.RegisterHandler(bus, handlers.OrderPlacedHandler)

	// Create events
	userEvents := []eventbus.Event[events.UserCreatedEvent]{
		{Data: events.UserCreatedEvent{UserName: "doe", UserID: 1}},
		{Data: events.UserCreatedEvent{UserName: "do", UserID: 2}}, // Should be skipped
		{Data: events.UserCreatedEvent{UserName: "john", UserID: 3}},
	}

	orderEvents := []eventbus.Event[events.OrderPlacedEvent]{
		{Data: events.OrderPlacedEvent{OrderID: 1, Amount: 10}},
		{Data: events.OrderPlacedEvent{OrderID: 2, Amount: -90}}, // Should be skipped
		{Data: events.OrderPlacedEvent{OrderID: 3, Amount: 1000}},
	}

	// Dispatch events
	for _, event := range userEvents {
		if err := eventbus.Dispatch(bus, event); err != nil {
			fmt.Printf("Error dispatching user event: %v\n", err)
		}
	}

	for _, event := range orderEvents {
		if err := eventbus.Dispatch(bus, event); err != nil {
			fmt.Printf("Error dispatching order event: %v\n", err)
		}
	}
}
