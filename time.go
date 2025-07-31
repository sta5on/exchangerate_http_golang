package main

import (
	"log"
	"time"
)

func convertTime(unix int) string {
	t := time.Unix(int64(unix), 0)
	return t.Format("15:04 MST")
}

func getCurrTime() string {
	currentUnix := int(time.Now().Unix())
	return convertTime(currentUnix)
}

func getTime(file string) (timeLast int) {
	rates, err := loadRates(file)
	if err != nil {
		log.Fatal(err)
	}
	timeLast = rates.TimeLastUpdateUnix

	return timeLast
}

func getActualityTime() (hours, minutes int) {
	currentUnix := int(time.Now().Unix())
	targetUnix := getTime("response.json")
	diffSeconds := currentUnix - targetUnix

	hours = diffSeconds / 3600
	minutes = (diffSeconds % 3600) / 60
	return hours, minutes
}
