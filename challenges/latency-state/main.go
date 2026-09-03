package latencystate

import "fmt"

func main() {
	var requestsNumber int
	var totalLatency float64

	fmt.Print("Number of requests: ")
	fmt.Scan(&requestsNumber)
	var latencyInput float64

	for rn := range requestsNumber {
		step := rn + 1
		fmt.Printf("Request %v latency: ", step)
		fmt.Scan(&latencyInput)
		totalLatency += latencyInput
	}

	avg := totalLatency / float64(requestsNumber)
	fmt.Printf("\nTotal latency: %.2f ms\n", totalLatency)
	fmt.Printf("Average latency: %.2f ms\n", avg)

	if avg <= 100 {
		fmt.Println("STATUS: FAST")
	} else if avg <= 300 {
		fmt.Println("STATUS: ACCEPTABLE")
	} else {
		fmt.Println("STATUS: SLOW")
	}
}