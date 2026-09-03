package main

import (
	"fmt"
)

func calculateAverage(readings []float64) float64 {
	numberReading := len(readings)
	totalReading := 0.0

	for _, value := range readings {
		totalReading += value
	}

	return totalReading / float64(numberReading)
}

func findHighest(readings []float64) float64 {
	highestNum := readings[0]

	for _, value := range readings {
		if highestNum < value {
			highestNum = value
		}
	}

	return highestNum
}

func countAboveAverage(readings []float64, avg float64) int {
	var counter int

	for _, value := range readings {
		if value > avg {
			counter++
		}
	}

	return counter
}

func main() {
	var numberOfSensors int
	var sensorVal float64

	sensors := []float64{}

	fmt.Print("Number for reading: ")
	fmt.Scan(&numberOfSensors)

	for num := range numberOfSensors {
		step := num + 1

		fmt.Printf("Reading %d: ", step)
		fmt.Scan(&sensorVal)
		sensors = append(sensors, sensorVal)
	}

	avg := calculateAverage(sensors)

	fmt.Printf("\nAverage: %.2f\n", avg)
	fmt.Printf("Highest: %.2f\n", findHighest(sensors))
	fmt.Printf("Above Average: %d\n", countAboveAverage(sensors, avg))
}
