package main

import "fmt"

func main() {
	var USD_EUR = 0.9
	var USD_RUB = 80.0
	var EUR_RUB = USD_RUB / USD_EUR
	fmt.Printf("Курс EUR к RUB %v", EUR_RUB)
}
