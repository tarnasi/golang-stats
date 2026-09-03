package main

import (
	"fmt"
)


func calculateDownloadTime(fileSizeMB float64, speedMBps float64) float64 {
	// size in MB * 8 / internet speed in Mbps
	return fileSizeMB * 8 / speedMBps
}

func classifyDownload(seconds float64) string {
	if seconds <= 60 {
		return "FAST"
	} else if seconds <= 300 {
		return "MODERATE"
	} else {
		return "SLOW"
	}
}


func main() {
	var fileSize float64
	var internetSpeed float64

	fmt.Print("File Size in MB: ")
	fmt.Scan(&fileSize)

	fmt.Print("Internet speed: ")
	fmt.Scan(&internetSpeed)

	resultInSeconds := calculateDownloadTime(fileSize, internetSpeed)
	fmt.Printf("\nDownload Time: %.2f seconds", resultInSeconds)
	fmt.Printf("\nStatus: %s\n", classifyDownload(resultInSeconds))
}


