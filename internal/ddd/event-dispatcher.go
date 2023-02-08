package ddd

import (
	"context"
	"sync"
)

type EventSubscriber interface {
	Subscribe(event Event, handler EventHandler)
}

type EventPublisher interface {
	Publish(ctx context.Context, events ...Event) error
}

type EventDispatcher struct {
	handlers map[string][]EventHandler
	mutex    sync.Mutex
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (dispatcher *EventDispatcher) Subscribe(event Event, handler EventHandler) {
	dispatcher.mutex.Lock()
	defer dispatcher.mutex.Unlock()

	dispatcher.handlers[event.GetName()] = append(dispatcher.handlers[event.GetName()], handler)
}

func (dispatcher *EventDispatcher) Publish(ctx context.Context, events ...Event) error {
	for _, event := range events {
		for _, handler := range dispatcher.handlers[event.GetName()] {
			err := handler(ctx, event)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
