package main

import (
	"os"
)

func main() {
	data1 := []byte("Hi\n Pond")
	err := os.WriteFile("data.txt", data1, 0644)
	if err != nil {
		panic(err)
	}
}
