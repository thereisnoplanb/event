package event

import "sync"

type Event[TSender any, TEventArgs any] struct {
	delegates map[uint32]EventHandler[TSender, TEventArgs]
	index     uint32
	mutex     *sync.Mutex
}

type Handler struct {
	index uint32
}

// Creates a new Event instance.
//
// # Returns
//
//	Event[TSender, TEventArgs]
//
// A new Event instance.
func New[TSender any, TEventArgs any]() Event[TSender, TEventArgs] {
	return Event[TSender, TEventArgs]{
		delegates: make(map[uint32]EventHandler[TSender, TEventArgs]),
		index:     0,
		mutex:     &sync.Mutex{},
	}
}

// Adds delegate function to event.
//
// # Parameters
//
//	delegate EventHandler[TSender, TEventArgs]
//
// Delegate function to add. If delegate is nil it is not added to event.
//
// # Returns
//
//	handler *Handler
//
// Returns handler to added delegate. If added delegate is nil, nil is returned.
func (event *Event[TSender, TEventArgs]) Add(delegate EventHandler[TSender, TEventArgs]) (handler *Handler) {
	if event != nil && delegate != nil {
		event.mutex.Lock()
		defer event.mutex.Unlock()
		event.index++
		event.delegates[event.index] = delegate
		return &Handler{
			index: event.index,
		}
	}
	return nil
}

// Removes delegate function from event.
//
// # Parameters
//
//	handler *Handler
//
// Handler to delegate, received in Add(delegate func(sender *TSender, eventArgs TEventArgs)) method.
func (event *Event[TSender, TEventArgs]) Remove(handler *Handler) {
	if event != nil && handler != nil {
		event.mutex.Lock()
		defer event.mutex.Unlock()
		delete(event.delegates, handler.index)
		handler.index = 0
	}
}

// Invokes event.
// Every added delegate is invoked in indeterminate order, not necessarily in the order of addition.
//
// # Parameters
//
//	sender *TSender
//
// Object that sends event.
//
//	eventArgs TEventArgs
//
// Arguments of event.
func (event *Event[TSender, TEventArgs]) Invoke(sender *TSender, eventArgs TEventArgs) {
	if event != nil {
		event.mutex.Lock()
		defer event.mutex.Unlock()
		for _, delegate := range event.delegates {
			delegate(sender, eventArgs)
		}
	}
}
