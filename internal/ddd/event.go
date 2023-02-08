package ddd

import (
	"context"
)

type Event interface {
	GetName() string
}

type EventHandler func(cxt context.Context, event Event) error
