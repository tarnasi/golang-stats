package tempertaturestatus

import "fmt"


func main() {
	var temperature float64

	fmt.Print("Temperature: ")
	fmt.Scan(&temperature)

	if temperature <= 60.0 {
		fmt.Println("STATUS: NORMAL")
	} else if temperature <= 80.0 {
		fmt.Println("STATUS: WARNING")
	} else {
		fmt.Println("STATUS: CRITICAL")
	}
}