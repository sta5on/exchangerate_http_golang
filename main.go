package main

import (
	"fmt"
	"log"
)

var option int
var curr1, curr2 string
var rate1, rate2 float64
var result float64
var amountstr string
var amount float64

func main() {
	log.Println("App is starting")
	welcomeMsg()
	defer log.Println("App is ending")
	// проверка жсон на актуальность
	go checkFileActual()
MainLoop:
	for {
		for {
			fmt.Printf("\nChoose your option:\n\n1 - Calculate exchange\n2 - Check rate\n3 - See all currencies\n4 - Exit\n")
			fmt.Scanln(&option)
			if option >= 1 && option <= 4 {
				break
			}
			fmt.Println("Invalid operation, please choose 1-4")
		}
		switch option {
		case 1:
			currInput()
			rate1, rate2 = getValues("response.json", curr1, curr2)
			amountInput()

			rate1, rate2, amount, result, curr1, curr2 = exchange(rate1, rate2, amount, curr1, curr2)
			fmt.Println(result)
			fmt.Printf("\nYou exchange %.2f %s TO %s\n%s/%s rate is %.2f\nFor %.2f %s YOU WILL GET %.2f %s\n\n", amount, curr1, curr2, curr1, curr2, rate2, amount, curr1, result, curr2)

			fmt.Printf("Rate is actual for: %s\n", convertTime(getTime("response.json")))

			//hr, mins := getActualityTime()
			//fmt.Printf("Differnce is %d hours and %d minutes\n", hr, mins)
			//fmt.Println(convertTime(getTime("response.json")))
			//fmt.Println("Current time:", getCurrTime())
		case 2:
			currInput()
			rate1, rate2 = getValues("response.json", curr1, curr2)
			if rate1 == 1 || rate1 == 1.0 {
				fmt.Printf("\nFirst Rate %d, curr: %v, second Rate %.4f, curr %s \n", int(rate1), curr1, rate2, curr2)
			} else if rate2 == 1 || rate2 == 1.0 {
				fmt.Printf("\nFirst Rate %.4f, curr: %v, second Rate %d, curr %s \n", rate1, curr1, int(rate2), curr2)
			} else {
				fmt.Printf("\nFirst Rate %.4f curr: %v, second Rate %.4f, curr %s \n", rate1, curr1, rate2, curr2)
			}
		case 3:
			showCurr()
		case 4:
			fmt.Println("Exiting...")
			break MainLoop
		}
	}
	fmt.Printf("\n0_0\n\n")
}
