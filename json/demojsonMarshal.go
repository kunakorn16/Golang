package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	data, _ := json.Marshal(&employee{101, "Kunakorn Klaewkla", "0983145470", "kunakorn.k@rmutsvmail.com"})
	fmt.Println(string(data))
}
