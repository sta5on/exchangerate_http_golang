package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
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

	curr1 := "EUR"
	curr2 := "MDL"
	result := 0.0

	rate1, rate2 := getValues("response.json", curr1, curr2)
	if rate1 == 1 || rate1 == 1.0 {
		fmt.Printf("\nFirst Rate %d, curr: %v, second Rate %.4f, curr %s \n", int(rate1), curr1, rate2, curr2)
	} else if rate2 == 1 || rate2 == 1.0 {
		fmt.Printf("\nFirst Rate %.4f, curr: %v, second Rate %d, curr %s \n", rate1, curr1, int(rate2), curr2)
	} else {
		fmt.Printf("\nFirst Rate %.4f curr: %v, second Rate %.4f, curr %s \n", rate1, curr1, rate2, curr2)
	}
	var amountstr string
	fmt.Print("Enter amount: ")

	_, err := fmt.Scanln(&amountstr)
	if err != nil {
		log.Println("Ошибка ввода:", err)
	}
	amount, err := strconv.ParseFloat(amountstr, 64)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Вы ввели: %f\n", amount)

	rate1, rate2, amount, result, curr1, curr2 = exchange(rate1, rate2, amount, curr1, curr2)
	fmt.Println(result)
	fmt.Printf("\nYou exchange %.f %s TO %s\n%s/%s rate is %f\nFor %f %s YOU WILL GET %f %s\nTHANK YOU!\n", amount, curr1, curr2, curr1, curr2, rate2, amount, curr1, result, curr2)
	// сделать по разному отображение дробной части
	// если ее нет то сократить нули если есть то показывать 2 символа наверное
	//
	//в вывод этого говна доабвить дату время актуальности курса! unix тайм взять как то перевести мб или +3 к гмт добавить к тому че имеем хз
	// сделать отдельными функциями getTime и вывод норм сделать
	//
}

func exchange(rate1, rate2, amount float64, curr1, curr2 string) (float64, float64, float64, float64, string, string) {
	//100 юсд в евро
	//100*0.86=86 EUR
	//100 евро в mdl
	//100/0.8681*17.02 = 1957

	result := 0.0
	if rate1 == 1.0 {
		fmt.Println("One of you currency is USD")
		result = amount * rate2
	}
	// сделать перевод из чего то другого кроме долларов и наоборот

	//100 mdl to usd = 100 mdl/17.02

	if rate2 == 1.0 {
		fmt.Println("usd is second")
		result = amount / rate1
	}

	// 100 eur to mdl
	// eur -> usd -> mdl

	result = amount / rate1 * rate2

	return rate1, rate2, amount, result, curr1, curr2
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

//func calculateReturn()

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
