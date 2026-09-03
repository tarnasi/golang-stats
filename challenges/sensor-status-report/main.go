package main

import "fmt"

type Sensor struct {
	Name     string
	Value    float64
	MinValue float64
	MaxValue float64
}

func classifySensor(sensor Sensor) string {
	if sensor.MinValue > sensor.MaxValue {
		return "INVALID"
	} else if sensor.Value < sensor.MinValue {
		return "LOW"
	} else if sensor.Value > sensor.MaxValue {
		return "HIGH"
	}

	return "NORMAL"
}

func main() {
	var numberOfSensors int
	sensors := []Sensor{}

	fmt.Print("Number of sensors: ")
	fmt.Scan(&numberOfSensors)

	for val := range numberOfSensors {
		var sensor Sensor

		fmt.Printf("\nSensor %d name: ", val+1)
		fmt.Scan(&sensor.Name)
		fmt.Print("Current Value: ")
		fmt.Scan(&sensor.Value)
		fmt.Print("Minimum Value: ")
		fmt.Scan(&sensor.MinValue)
		fmt.Print("Maximum Value: ")
		fmt.Scan(&sensor.MaxValue)
		sensors = append(sensors, sensor)
	}

	fmt.Print("\n--- Sensor Report ---")
	for _, sensor := range sensors {
		status := classifySensor(sensor)
		fmt.Printf("%s: %.2f - %s\n", sensor.Name, sensor.Value, status)
	}
}
