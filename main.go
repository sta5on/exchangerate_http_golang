package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
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

	fmt.Println("Сырой JSON:")
	fmt.Println(string(body))

	//json.Unmarshal()
}
