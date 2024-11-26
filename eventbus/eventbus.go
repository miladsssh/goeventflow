package eventbus

import (
	"fmt"
	"reflect"
	"sync"
)

type Event[T any] struct {
	Data T
}

type EventBus struct {
	handlers          sync.Map
	middlewares       sync.Map
	globalMiddlewares []GlobalMiddlewareFunc
}

func Dispatch[T any](bus *EventBus, event Event[T]) error {
	eventType := reflect.TypeOf(event.Data)

	// Process global middlewares
	for _, middleware := range bus.globalMiddlewares {
		proceed, err := middleware(event.Data)
		if err != nil {
			return fmt.Errorf("global middleware error: %w", err)
		}
		if !proceed {
			return nil
		}
	}

	// Process type-specific middlewares
	if value, ok := bus.middlewares.Load(eventType); ok {
		middlewareList := *(value.(*[]MiddlewareFunc[T]))
		for _, middleware := range middlewareList {
			proceed, err := middleware(&event)
			if err != nil {
				return fmt.Errorf("middleware error: %w", err)
			}
			if !proceed {
				return nil
			}
		}
	}

	// Invoke handlers
	if value, ok := bus.handlers.Load(eventType); ok {
		handlerList := *(value.(*[]EventHandlerFunc[T]))
		var wg sync.WaitGroup
		for _, handler := range handlerList {
			wg.Add(1)
			go func(handler EventHandlerFunc[T]) {
				defer wg.Done()
				if err := handler(event); err != nil {
					fmt.Printf("Handler error: %v\n", err)
				}
			}(handler)
		}
		wg.Wait()
	}

	return nil
}

func NewEventBus() *EventBus {
	return &EventBus{}
}
