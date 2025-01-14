package event

type EventHandler[TSender any, TEventArgs any] func(sender *TSender, eventArgs TEventArgs)
