package eventbus

import "reflect"

type MiddlewareFunc[T any] func(event *Event[T]) (proceed bool, err error)

type GlobalMiddlewareFunc func(event any) (proceed bool, err error)

func (bus *EventBus) RegisterGlobalMiddlewares(middlewares ...GlobalMiddlewareFunc) {
	bus.globalMiddlewares = append(bus.globalMiddlewares, middlewares...)
}

func RegisterMiddleware[T any](bus *EventBus, middlewares ...MiddlewareFunc[T]) {
	eventType := reflect.TypeOf((*T)(nil)).Elem()
	value, _ := bus.middlewares.LoadOrStore(eventType, &[]MiddlewareFunc[T]{})
	middlewareList := value.(*[]MiddlewareFunc[T])
	*middlewareList = append(*middlewareList, middlewares...)
}
