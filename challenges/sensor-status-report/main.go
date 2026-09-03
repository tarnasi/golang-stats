package main

import "fmt"

type Sensor struct {
	Name     string
	Value    float64
	MinValue float64
	MaxValue float64
}

type ReportSensor struct {
	Name   string
	Value  float64
	Status string
}

var stores []Sensor
var reports []ReportSensor

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
	var numberOfSensor int

	fmt.Print("Number of sensors: ")
	fmt.Scan(&numberOfSensor)

	for val := range numberOfSensor {
		var sensor Sensor
		var report ReportSensor

		fmt.Printf("\nSensor %d name: ", val+1)
		fmt.Scan(&sensor.Name)
		fmt.Print("Current Value: ")
		fmt.Scan(&sensor.Value)
		fmt.Print("Minimum Value: ")
		fmt.Scan(&sensor.MinValue)
		fmt.Print("Maximum Value: ")
		fmt.Scan(&sensor.MaxValue)

		stores = append(stores, sensor)
		status := classifySensor(sensor)

		report.Name = sensor.Name
		report.Value = sensor.Value
		report.Status = status

		reports = append(reports, report)
	}

	fmt.Println("--- Sensor Report ---")
	for _, report := range reports {
		fmt.Printf("\n%v: %.2f - %v", report.Name, report.Value, report.Status)
	}
}
