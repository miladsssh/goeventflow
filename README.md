# GoEventFlow

GoEventFlow is a generic, event-driven middleware system in Go that supports:

- **Dynamic Event Registration**: Register handlers and middleware for specific event types at runtime.
- **Type-Safe Event Handling**: Utilize Go Generics to ensure that handlers receive events of the correct type.
- **Middleware Support**: Process events through middleware before they reach handlers.
- **Concurrent Event Dispatching**: Dispatch events concurrently, with thread-safe handler execution.
- **Global Middlewares**: Apply middleware functions to all event types.

## Features

- **Generic Event Types**: Define events with any data payload.
- **Flexible Middleware**: Middleware can modify events or halt propagation based on conditions.
- **Custom Serialization**: Events can be serialized/deserialized with custom JSON marshaling.
- **Error Handling**: Errors in middleware and handlers are handled gracefully.