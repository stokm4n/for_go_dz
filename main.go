package main

import "fmt"

func main() {
	// const USDtoEURO float64 = 0.86
	// const USDtoRUB float64 = 77.89
	// const EUROtoRUB float64 = USDtoRUB / USDtoEURO
	currency1, currency2, quantity := getUserInput()
	newQuantity := calculateCurrency(currency1, currency2, quantity)
	fmt.Print(newQuantity)
}

func getUserInput() (string, string, float64) {
	var currency1, currency2 string
	var quantity float64
	fmt.Println("Введите название исходной валюты: ")
	fmt.Scan(&currency1)
	fmt.Println("Введите название целевой валюты: ")
	fmt.Scan(&currency2)
	fmt.Print("Введите количество исходной валюты: ")
	fmt.Scan(&quantity)
	return currency1, currency2, quantity
}

func calculateCurrency(currency1 string, currency2 string, quantity float64) float64 {
	return 0 // Пустая функция, реализация будет в следующем коммите
}
