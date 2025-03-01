# Event Library

The event library in Go allows for managing events and delegates. Below is a description of the functions and how to use this library.

## Installation

To install the library, use the following command:
```sh
go get github.com/thereisnoplanb/event
```

## Usage
### Creating a New Event
```golang
package main
    
import (
    "fmt"
    "github.com/your-username/event"
)

type Test struct {
    Event event.Event[Test, EventArgsTest]
}

func New() Test {
    
}

func main() {
    event := event.New[string, string]()

    handler := event.Add(func(sender string, eventArgs string) {
        fmt.Println("Event received from:", sender, "with args:", eventArgs)
    })

    event.Invoke("Sender", "Hello, World!")
    event.Remove(handler)
}
```

#### Event Interface
The Event interface defines methods for managing events:
- Add(delegate EventHandler[TSender, TEventArgs]) (handler Handler): Adds a delegate function to the event.
- Remove(handler *Handler): Removes a delegate function from the event.
- Invoke(sender *TSender, eventArgs TEventArgs): Invokes the event.

#### Handler Struct
The Handler struct stores the index of the delegate added to the event.

#### New Function
The New function creates a new event instance:
```
func New[TSender any, TEventArgs any]() Event[TSender, TEventArgs]
```
##### Example
Below is a complete example of using the library:
```go
package main

import (
"fmt"
"github.com/your-username/event"
)

func main() {
event := event.New[string, string]()

handler := event.Add(func(sender string, eventArgs string) {
fmt.Println("Event received from:", sender, "with args:", eventArgs)
})

event.Invoke("Sender", "Hello, World!")
event.Remove(handler)
}
```

# License
This project is licensed under the terms of the MIT license.
See the LICENSE file for more information.


