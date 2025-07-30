package main

import (
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

	response := string(body)
	write_to_file(response)
	//json.Unmarshal()
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
