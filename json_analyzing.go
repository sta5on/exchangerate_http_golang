package main

import (
	"log"
	"os"
)

// func sss()  {
//
// }
func govno() {
	filePath := "example.json"
	data := []byte("Hello, Go!\nThis is a new line.")

	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Data written to", filePath)
}

func main() {
	govno()
}
