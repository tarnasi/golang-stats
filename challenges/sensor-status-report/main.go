package main

import "fmt"

type Sensor struct {
	Name     string
	Value    float64
	MinValue float64
	MaxValue float64
}

func (sensor Sensor) Status() string {
	if !sensor.IsValid() {
		return "INVALID"
	} else if sensor.Value < sensor.MinValue {
		return "LOW"
	} else if sensor.Value > sensor.MaxValue {
		return "HIGH"
	}

	return "NORMAL"
}

func (sensor Sensor) IsValid() bool {
    return sensor.MinValue <= sensor.MaxValue
}

func (sensor Sensor) Deviation() float64 {
	status := sensor.Status()

	if !sensor.IsValid() || status == "NORMAL" {
		return 0
	} else if status == "LOW" {
		return sensor.MinValue - sensor.Value
	}

	// High
	return sensor.Value - sensor.MaxValue
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

	fmt.Println("\n--- Sensor Report ---")
	for _, sensor := range sensors {
		status := sensor.Status()
		deviation := sensor.Deviation()
		fmt.Printf("%s: %.2f - %s - Deviation: %.2f\n", sensor.Name, sensor.Value, status, deviation)
	}
}
