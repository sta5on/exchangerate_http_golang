package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

func currInput() {
	for {
		fmt.Print("Enter first currency: ")
		_, err := fmt.Scanln(&curr1)
		if err != nil {
			log.Println("Ошибка ввода:", err)
			continue
		}

		curr1 = strings.ToUpper(curr1)
		if slices.Contains(currencyList[:], curr1) {
			break
		}
		fmt.Printf("%s not found! Try again.\n", curr1)
	}

	for {
		fmt.Print("Enter second currency: ")
		_, err := fmt.Scanln(&curr2)
		if err != nil {
			log.Println("Ошибка ввода:", err)
			continue
		}

		curr2 = strings.ToUpper(curr2)
		if curr1 != curr2 {
			if slices.Contains(currencyList[:], curr2) {
				break
			}
		}
		if curr1 == curr2 {
			fmt.Println("You can`t use 2 same currencies")
		} else {
			fmt.Printf("%s not found! Try again.\n", curr2)
		}
	}

	fmt.Printf("Вы ввели: %s и %s\n", curr1, curr2)
}

func amountInput() {
	for {
		fmt.Print("Enter amount: ")

		_, err := fmt.Scanln(&amountstr)
		if err != nil {
			log.Println("Ошибка ввода:", err)
			continue
		}
		amount, err = strconv.ParseFloat(amountstr, 64)
		if err != nil {
			log.Println(err)
			continue
		}
		if err == nil {
			break
		}
		fmt.Println("Некорректный ввод")
	}
	fmt.Printf("Вы ввели: %f\n", amount)
}
