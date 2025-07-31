package main

import (
	"log"
	"time"
)

// Запрос по апи, сохранение ответа в жсон response.json
//
// writeToFile(getRates())

func checkFileActual() {
	timeLast := getTime("response.json")
	timeNow := int(time.Now().Unix())
	if timeNow >= (timeLast + 3600) {
		updateFile()
	}
	// 1 hr is 3600 sec
}

func updateFile() {
	for {
		writeToFile(getRates())
		hr, mins := getActualityTime()
		log.Printf("Differnce is %d hours and %d minutes\n", hr, mins)
		log.Println(convertTime(getTime("response.json")))
		log.Println("Current time:", getCurrTime())
		time.Sleep(1 * time.Hour)
	}
}

//go checkFileActual
