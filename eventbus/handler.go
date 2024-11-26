package eventbus

import (
	"reflect"
)

type EventHandlerFunc[T any] func(event Event[T]) error

func RegisterHandler[T any](bus *EventBus, handlers ...EventHandlerFunc[T]) {
	eventType := reflect.TypeOf((*T)(nil)).Elem()
	value, _ := bus.handlers.LoadOrStore(eventType, &[]EventHandlerFunc[T]{})
	handlerList := value.(*[]EventHandlerFunc[T])
	*handlerList = append(*handlerList, handlers...)
}
