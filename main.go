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
	defer log.Println("\nApp is ending")
	//
	//Запрос по апи, сохранение ответа в жсон response.json
	//
	//write_to_file(getRates())
	//

	curr1 := "USD"
	curr2 := "EUR"

	rate1, rate2 := getValues("response.json", curr1, curr2)
	fmt.Printf("\nFirst Rate %s, curr: %v, second Rate %f, curr %s \n", rate1, curr1, rate2, curr2)

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

func getValues(file string, curr1 string, curr2 string) (rate1 float64, rate2 float64) {
	rates, err := loadRates(file)
	if err != nil {
		log.Fatal(err)
	}
	rate1, ok := rates.ConversionRates[curr1]
	if !ok {
		fmt.Println("Currency &s not found", curr2)
	}
	rate2, ok = rates.ConversionRates[curr2]
	if !ok {
		fmt.Println("Currency &s not found", curr2)
	}
	return rate1, rate2
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
