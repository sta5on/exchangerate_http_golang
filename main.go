package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	api_key = "3e6087d3800124a5a3cb77e7"
)

// https://v6.exchangerate-api.com/v6/3e6087d3800124a5a3cb77e7/latest/USD
func GetExchangeURL(currency string) string {
	return fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/latest/%s", api_key, currency)
}

func main() {
	log.Println("App is starting")
	//
	//Запрос по апи, сохранение ответа в жсон response.json
	//
	//write_to_file(getRates())
	//
	defer log.Println("App is ending")
	////json.Unmarshal()
	fmt.Println(loadRates("response.json"))
}

func getRates() (response string) {
	resp, err := http.Get(GetExchangeURL("USD"))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Ошибка чтения ответа:", err)
	}

	response = string(body)
	return response
}

func write_to_file(response string) {
	filePath := "response.json"
	data := []byte(response)

	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Data written to", filePath)
}

func loadRates(file string) (*response_usd, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var rates response_usd

	err = json.Unmarshal(data, &rates)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &rates, nil
}
