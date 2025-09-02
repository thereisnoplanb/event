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
    "github.com/thereisnoplanb/event"
)

func main() {
    event, invoke := event.New[string, string]()

    handler := event.Add(onInvoke)

    event.Invoke("Sender", "Hello, World!")
    event.Remove(handler)
}

func onInvoke(sender string, eventArgs string) {
    fmt.Println("Event received from:", sender, "with args:", eventArgs)
}
```

#### Event Interface
The Event interface defines methods for managing events:
- Add(delegate EventHandler[TSender, TEventArgs]) (handle *Handle): Adds a delegate function to the event.
- Remove(handle *Handle): Removes a delegate function from the event.

#### Inoke method
The method to
- Invoke(sender *TSender, eventArgs TEventArgs): Invokes the event.

#### Handler Struct
The Handler struct stores the UUID of the delegate added to the event.

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

	"github.com/thereisnoplanb/event"
)

type Thermometer struct {
	temperature              float64
	TemperatureChanged       event.Event[Thermometer, TemperatureChangedEventArgs]
	invokeTemperatureChanged func(sender *Thermometer, eventArgs TemperatureChangedEventArgs)
}

type TemperatureChangedEventArgs struct {
	PreviousTemperature float64
	ActualTemperature   float64
}

func New(temperature float64) (thermometer *Thermometer) {
	thermometer = &Thermometer{
		temperature: temperature,
	}
	thermometer.TemperatureChanged, thermometer.invokeTemperatureChanged = event.New[Thermometer, TemperatureChangedEventArgs]()
	return thermometer
}

func (thermometer *Thermometer) ChangeTemperature(temperature float64) {
	previousTemperature := thermometer.temperature
	thermometer.temperature = temperature
	thermometer.invokeTemperatureChanged(thermometer, TemperatureChangedEventArgs{
		PreviousTemperature: previousTemperature,
		ActualTemperature:   temperature,
	})
}

func main() {

	thermometer := New(0)
	handler := thermometer.TemperatureChanged.Add(onTemperatureChanged)
	defer thermometer.TemperatureChanged.Remove(handler)

	thermometer.ChangeTemperature(6)
	thermometer.ChangeTemperature(10)
}

func onTemperatureChanged(sender *Thermometer, eventArgs TemperatureChangedEventArgs) {
	fmt.Printf("Temperature has changed from %.1f to %.1f\n", eventArgs.PreviousTemperature, eventArgs.ActualTemperature)
}
```

#### Result
```sh
Temperature has changed from 0.0 to 6.0
Temperature has changed from 6.0 to 10.0
```

# License
This project is licensed under the terms of the MIT license.
See the LICENSE file for more information.


