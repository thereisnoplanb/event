package event

import "sync"

type event[TSender any, TEventArgs any] struct {
	delegates map[uint32]EventHandler[TSender, TEventArgs]
	index     uint32
	mutex     *sync.Mutex
}

type Handler struct {
	index uint32
}

type Event[TSender any, TEventArgs any] interface {
	// Adds delegate function to event.
	//
	// Parameters:
	//
	//	delegate func(sender *TSender, eventArgs TEventArgs) - Delegate function to add. If delegate is nil it is not added to event.
	//
	// Returns:
	//
	//	handler *Handler - Returns handler to added delegate. If added delegate is nil, nil is returned.
	Add(delegate func(sender *TSender, eventArgs TEventArgs)) (handler *Handler)

	// Removes delegate function from event.
	//
	// Parameters:
	//
	//	handler *Handler - Handler to delegate, received in Add(delegate func(sender *TSender, eventArgs TEventArgs)) method.
	Remove(handler *Handler)

	// Invokes event.
	// Every added delegate is invoked in indeterminate order, not necessarily in the order of addition.
	//
	// Parameters:
	//
	//	sender *TSender - Object that sends event.
	//	eventArgs TEventArgs - Arguments of event.
	Invoke(sender *TSender, eventArgs TEventArgs)
}

// Creates a new Event instance.
//
// Returns:
//
//	Event[TSender, TEventArgs] - A new Event instance.
func New[TSender any, TEventArgs any]() Event[TSender, TEventArgs] {
	return &event[TSender, TEventArgs]{
		delegates: make(map[uint32]EventHandler[TSender, TEventArgs]),
		index:     0,
		mutex:     &sync.Mutex{},
	}
}

// Adds delegate function to event.
//
// Parameters:
//
//	delegate func(sender *TSender, eventArgs TEventArgs) - Delegate function to add. If delegate is nil it is not added to event.
//
// Returns:
//
//	handler *Handler - Returns handler to added delegate. If added delegate is nil, nil is returned.
func (event *event[TSender, TEventArgs]) Add(delegate func(sender *TSender, eventArgs TEventArgs)) (handler *Handler) {
	if delegate != nil {
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
// Parameters:
//
//	handler *Handler - Handler to delegate, received in Add(delegate func(sender *TSender, eventArgs TEventArgs)) method.
func (event *event[TSender, TEventArgs]) Remove(handler *Handler) {
	if handler != nil {
		event.mutex.Lock()
		defer event.mutex.Unlock()
		delete(event.delegates, handler.index)
	}
}

// Invokes event.
// Every added delegate is invoked in indeterminate order, not necessarily in the order of addition.
//
// Parameters:
//
//	sender *TSender - Object that sends event.
//	eventArgs TEventArgs - Arguments of event.
func (event event[TSender, TEventArgs]) Invoke(sender *TSender, eventArgs TEventArgs) {
	event.mutex.Lock()
	defer event.mutex.Unlock()
	for _, delegate := range event.delegates {
		delegate(sender, eventArgs)
	}
}
