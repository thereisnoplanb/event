package event

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	type TemperatureChangedEventArgs struct {
		PreviousTemperature float64
		ActualTemperature   float64
	}
	type Thermometer struct {
		temperature             float64
		TemperatureChanged      Event[Thermometer, TemperatureChangedEventArgs]
		ChangeTempeartaure      func(thermometer *Thermometer, temperature float64)
		raiseTemperatureChanged func(thermometer *Thermometer, eventArgs TemperatureChangedEventArgs)
	}

	thermometer := &Thermometer{
		temperature: 0,
	}
	thermometer.TemperatureChanged, thermometer.raiseTemperatureChanged = New[Thermometer, TemperatureChangedEventArgs]()

	thermometer.ChangeTempeartaure = func(thermometer *Thermometer, temperature float64) {
		previousTempearature := thermometer.temperature
		thermometer.temperature = temperature
		if invokeEvent := thermometer.raiseTemperatureChanged; invokeEvent != nil {
			invokeEvent(thermometer, TemperatureChangedEventArgs{
				PreviousTemperature: previousTempearature,
				ActualTemperature:   temperature,
			})
		}
	}

	raised := false
	thermometer.TemperatureChanged.Add(func(sender *Thermometer, eventArgs TemperatureChangedEventArgs) {
		fmt.Printf("Temperature has changed from %.1f to %.1f\n", eventArgs.PreviousTemperature, eventArgs.ActualTemperature)
		raised = true
	})

	thermometer.ChangeTempeartaure(thermometer, 5)

	if !raised {
		t.Errorf("Event has not been raised")
	}
}
