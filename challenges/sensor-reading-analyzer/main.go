package main


import (
	"fmt"
)

func calculateAverage(reading []float64) float64 {
	numberReading := len(reading)
	totalReading := 0.0

	for _, value := range reading {
		totalReading += value
	}

	return totalReading / float64(numberReading)
}

func findHighest(reading []float64) float64 {
	highestNum := reading[0]

	for _, value := range reading {
		if highestNum < value {
			highestNum = value
		}
	}

	return highestNum
}

func countAboveAverage(reading []float64, avg float64) int {
	var counter int

	for _, value := range reading {
		if value > avg {
			counter += 1
		}
	}

	return counter
}

func main() {
	var number_of_sensors int
	var sensor_val float64

	sensors := []float64{}

	fmt.Print("Number for reading: ")
	fmt.Scan(&number_of_sensors)

	for num := range number_of_sensors {
		step := num + 1

		fmt.Printf("Reading %d: ", step)
		fmt.Scan(&sensor_val)
		sensors = append(sensors, sensor_val)
	}

	avg := calculateAverage(sensors)

	fmt.Printf("\nAverage: %.2f\n", avg)
	fmt.Printf("Above: %.2f\n", findHighest(sensors))
	fmt.Printf("Above Average: %d\n", countAboveAverage(sensors, avg))
}
