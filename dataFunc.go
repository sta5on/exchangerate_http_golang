package main

import (
	"fmt"
)

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
	//
	// 1 usd
	//0.8 eur
	//17 mdl
	//1/0.8*17

	result = amount / rate1 * rate2
	// 1 / 0.8 * 17
	rate2 = 1 / rate1 * rate2

	return rate1, rate2, amount, result, curr1, curr2
}
