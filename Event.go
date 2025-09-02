package event

import (
	"sync"

	"github.com/google/uuid"
)

type Handler[TSender any, TEventArgs any] func(sender *TSender, eventArgs TEventArgs)
type Invoke[TSender any, TEventArgs any] func(sender *TSender, eventArgs TEventArgs)

type Event[TSender any, TEventArgs any] interface {
	// Adds a handler function to the event.
	//
	// # Parameters
	//
	//	handler func(sender *TSender, eventArgs TEventArgs)
	//
	// The handler function to add.
	//
	// # Returns
	//
	//	handle *Handle
	//
	// Returns the handle to the added handler function.
	// Handle can be used to remove the handler function from the event.
	// See the Remove() method.
	Add(handler Handler[TSender, TEventArgs]) (handle *Handle)

	// Removes the handler function from the event.
	//
	// # Parameters
	//
	//	handle *Handle
	//
	// Handle to the handler function received in the Add() method.
	Remove(handle *Handle)
}

type event[TSender any, TEventArgs any] struct {
	handlers map[uuid.UUID]Handler[TSender, TEventArgs]
	mutex    *sync.Mutex
}

type Handle struct {
	uuid uuid.UUID
}

type EmptyEventArgs struct{}

// Creates a new Event instance.
//
// # Returns
//
//	event Event[TSender, TEventArgs]
//
// A new Event instance.
//
//	invoke Invoke[TSender, TEventArgs]
//
// A function to invoke the event.
func New[TSender any, TEventArgs any]() (_event Event[TSender, TEventArgs], invoke Invoke[TSender, TEventArgs]) {
	event := &event[TSender, TEventArgs]{
		handlers: make(map[uuid.UUID]Handler[TSender, TEventArgs]),
		mutex:    &sync.Mutex{},
	}
	return event, event.invoke
}

// Adds a handler function to the event.
//
// # Parameters
//
//	handler func(sender *TSender, eventArgs TEventArgs)
//
// The handler function to add.
//
// # Returns
//
//	handle *Handle
//
// Returns the handle to the added handler function.
// Handle can be used to remove the handler function from the event.
// See the Remove() method.
func (event *event[TSender, TEventArgs]) Add(handler Handler[TSender, TEventArgs]) (handle *Handle) {
	if handler == nil {
		return nil
	}
	event.mutex.Lock()
	defer event.mutex.Unlock()
	uuid := uuid.New()
	event.handlers[uuid] = handler
	return &Handle{
		uuid: uuid,
	}
}

// Removes the handler function from the event.
//
// # Parameters
//
//	handle *Handle
//
// Handle to the handler function. Received in the Add() method.
func (event *event[TSender, TEventArgs]) Remove(handle *Handle) {
	if handle != nil && handle.uuid != uuid.Nil {
		event.mutex.Lock()
		defer event.mutex.Unlock()
		delete(event.handlers, handle.uuid)
		handle.uuid = uuid.Nil
	}
}

// Invokes the event.
// Any non-nil handler function added is called in an unspecified order, not necessarily in the order they were added.
//
// # Parameters
//
//	sender *TSender
//
// A pointer to the event invoker.
//
//	eventArgs TEventArgs
//
// A TEventArgs that contains the event data.
func (event *event[TSender, TEventArgs]) invoke(sender *TSender, eventArgs TEventArgs) {
	event.mutex.Lock()
	defer event.mutex.Unlock()
	for _, handler := range event.handlers {
		if handler != nil {
			handler(sender, eventArgs)
		}
	}
}
