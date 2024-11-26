package middlewares

import (
	"fmt"

	"github.com/miladsssh/goeventflow/eventbus"
)

func LoggingMiddleware[T any](event *eventbus.Event[T]) (bool, error) {
	fmt.Printf("LoggingMiddleware - Event received: %+v\n", event.Data)
	return true, nil
}

func GlobalLoggingMiddleware(event any) (bool, error) {
	fmt.Printf("GlobalLoggingMiddleware - Event received: %+v\n", event)
	return true, nil
}
