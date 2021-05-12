package event

import (
	"context"
)

// Event is an abstraction for all messages that
// are sent to queue or received from queue.
type Event struct {
	// Key sets the key of the message for routing policy
	Key string
	// Payload for the message
	Payload []byte
	// Properties attach application defined properties on the message
	Properties map[string]string
}

// Handler is a callback function that processes messages delivered
// to asynchronous subscribers.
type Handler func(context.Context, Event) error

// Publisher is abstraction for sending messages
// to queue.
type Publisher interface {
	Publish(ctx context.Context, event Event) error
	Close() error
}

// Subscriber is an abstraction for receiving messages
// from queue.
type Subscriber interface {
	Subscribe(ctx context.Context, h Handler) error
	Close() error
}
