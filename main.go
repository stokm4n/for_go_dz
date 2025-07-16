package main

import (
	"fmt"
)

const USDTtoEURO float64 = 0.86
const USDTtoRUB float64 = 77.89
const RUBtoUSDT float64 = 0.013
const RUBtoEURO float64 = 0.011
const EUROtoRUB float64 = 90.9
const EUROtoUSDT float64 = 1.16

func main() {
	originalCurrency := getOriginalCurrency()
	value := getValue()
	targetCurrency := getTargetCurrency(originalCurrency)
	convertValue := calculateCurrency(originalCurrency, value, targetCurrency)
	fmt.Print("При конвертации ", value, " ", originalCurrency, " в ", targetCurrency, " получим: ", convertValue, " ", targetCurrency)
}

func getOriginalCurrency() string {
	var originalCurrency string
	fmt.Println("Введите исходную валюту. Доступные валюты: RUB, USDT, EURO. Ввод производить строго заглавными буквами")
	for {
		_, _ = fmt.Scan(&originalCurrency)
		if originalCurrency == "RUB" || originalCurrency == "USDT" || originalCurrency == "EURO" {
			return originalCurrency
		}
		fmt.Println("Вы ввели недоступную валюту или несуществующую валюту")
		fmt.Print("Введите еще раз исходную валюту: ")
	}
}

func getValue() float64 {
	var value float64
	for {
		fmt.Print("Введите количество исходной валюты: ")
		_, err := fmt.Scan(&value)
		if err != nil {
			fmt.Println("Ошибка: введено не число! Введите количество заново")
			var discard string
			fmt.Scanln(&discard)
		} else {
			return value
		}
	}
}

func getTargetCurrency(originalCurrency string) string {
	var targetCurrency string
	for {
		if originalCurrency == "RUB" {
			fmt.Println("Введите целевую валюту. Доступные валюты: USDT, EURO. Ввод производить строго заглавными буквами")
			_, _ = fmt.Scan(&targetCurrency)
			if targetCurrency == "USDT" || targetCurrency == "EURO" {
				return targetCurrency
			}
			fmt.Println("Вы ввели недоступную целевую валюту или несуществующую целевую валюту")
			fmt.Print("Введите еще раз целевую валюту: ")
		} else if originalCurrency == "USDT" {
			fmt.Println("Введите целевую валюту. Доступные валюты: RUB, EURO. Ввод производить строго заглавными буквами")
			_, _ = fmt.Scan(&targetCurrency)
			if targetCurrency == "RUB" || targetCurrency == "EURO" {
				return targetCurrency
			}
			fmt.Println("Вы ввели недоступную целевую валюту или несуществующую целевую валюту")
			fmt.Print("Введите еще раз целевую валюту: ")
		} else {
			fmt.Println("Введите исходную валюту. Доступные валюты: RUB, USDT. Ввод производить строго заглавными буквами")
			_, _ = fmt.Scan(&targetCurrency)
			if targetCurrency == "RUB" || targetCurrency == "USDT" {
				return targetCurrency
			}
			fmt.Println("Вы ввели недоступную целевую валюту или несуществующую целевую валюту")
			fmt.Print("Введите еще раз целевую валюту: ")
		}
	}
}

func calculateCurrency(originalCurrency string, value float64, targetCurrency string) float64 {
	if originalCurrency == "USDT" && targetCurrency == "EURO" {
		return USDTtoEURO * value
	} else if originalCurrency == "USDT" && targetCurrency == "RUB" {
		return USDTtoRUB * value
	} else if originalCurrency == "RUB" && targetCurrency == "USDT" {
		return RUBtoUSDT * value
	} else if originalCurrency == "RUB" && targetCurrency == "EURO" {
		return RUBtoEURO * value
	} else if originalCurrency == "EURO" && targetCurrency == "RUB" {
		return EUROtoRUB * value
	} else {
		return EUROtoUSDT * value
	}
}
