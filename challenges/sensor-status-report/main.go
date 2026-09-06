package main

import "fmt"

type Sensor struct {
	Name     string
	Value    float64
	MinValue float64
	MaxValue float64
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
		sensor, err := createSensor(
			sensor.Name,
			sensor.Value,
			sensor.MinValue,
			sensor.MaxValue,
		)

		if err != nil {
			fmt.Println("Error:", err)
    		continue
		}
		sensors = append(sensors, sensor)	
	}

	printReport(sensors, "Sensor Report Before Calibration")

	// Calibrate sensors
	for index := range sensors {
		var offset float64
		fmt.Printf("\nCalibration offset for sensor %s: ", sensors[index].Name)
		fmt.Scan(&offset)
		sensors[index].Calibrate(offset)
	}

	printReport(sensors, "Sensor Report After Calibration")
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

func printReport(sensors []Sensor, title string) {
	fmt.Printf("\n--- %s ---\n", title)
	for _, sensor := range sensors {
		status := sensor.Status()
		deviation := sensor.Deviation()
		fmt.Printf(
			"%s: %.2f - %s - Deviation: %.2f\n",
			sensor.Name,
			sensor.Value,
			status,
			deviation,
		)
	}
}

func (sensor *Sensor) Calibrate(offset float64) {
	sensor.Value += offset
}

func createSensor(
	name string,
	value float64,
	minValue float64,
	MaxValue float64,
) (Sensor, error) {
	sensor := Sensor {
		name, value, minValue, MaxValue,
	}

	if !sensor.IsValid() {
		return Sensor{}, fmt.Errorf(
			"minimum %.2f cannot be greater than maximum %.2f",
			minValue,
			MaxValue,
		)
	}

	return sensor, nil
}