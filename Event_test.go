package event

import (
	"fmt"
	"testing"
)

type Test struct {
	Event Event[Test, EventArgsTest]
}

func NewTest() Test {
	return Test{
		Event: New[Test, EventArgsTest](),
	}
}

func (t *Test) Do(message string) {
	t.Event.Invoke(t, EventArgsTest{
		AdditionalMessage: message,
	})
}

type EventArgsTest struct {
	AdditionalMessage string
}

func TestEvent(t *testing.T) {

	delegate1 := func(sender *Test, eventArgs EventArgsTest) {
		fmt.Printf("Sender: %v delegate 1 %s.\n", sender, eventArgs.AdditionalMessage)
	}
	delegate2 := func(sender *Test, eventArgs EventArgsTest) {
		fmt.Printf("Sender: %v delegate 2 %s.\n", sender, eventArgs.AdditionalMessage)
	}
	test := NewTest()
	id1 := test.Event.Add(delegate1)
	_ = test.Event.Add(delegate2)
	test.Do("Message 1")
	test.Event.Remove(id1)
	test.Do("Message 2")
}
