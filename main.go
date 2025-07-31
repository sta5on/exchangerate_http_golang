package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	log.Println("App is starting")
	welcomeMsg()
	defer log.Println("App is ending")
	go checkFileActual()
	//
	//Запрос по апи, сохранение ответа в жсон response.json
	//
	//writeToFile(getRates())
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

	// если НЕ ДОЛЛАР ТО EUR/MDL rate is 17.024895 строка неверна! исправить!

	fmt.Printf("Rate is actual for: %d\n", getTime("response.json"))

	hr, mins := getActualityTime()
	fmt.Printf("Differnce is %d hours and %d minutes\n", hr, mins)
	fmt.Println(convertTime(getTime("response.json")))
	fmt.Println("Current time:", getCurrTime())

	fmt.Printf("\n\n0_0\n\n")
}
