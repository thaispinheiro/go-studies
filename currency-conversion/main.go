package main

import (
	"fmt"
	"strconv"
	"strings"
)

var exchangeRates = map[string]float64 {
	"USD": 0.151,
	"EUR": 0.137,
	"JPY": 16.29,
	"GBP": 0.13,
	"CHF": 0.1402,
	"AUD": 0.2712,
	"CAD": 0.2374,
	"CNY": 1.251,
	"HKD": 1.326,
	"NZD": 0.2922,
	"SEK": 1.655,
	"NOK": 1.806,
	"DKK": 1.122,
	"SGD": 0.2249,
	"KRW": 242.97,
	"ZAR": 3.239,
	"MXN": 3.454,
	"INR": 14.71,
	"ILS": 0.63,
	"THB": 5.74,
	"IDR": 2875.0,
	"MYR": 0.754,
	"PHP": 9.74,
	"PLN": 0.644,
	"CZK": 3.77,
	"HUF": 61.59,
	"TRY": 6.49,
	"BGN": 0.293,
	"RON": 0.746,
}

func main() {
	targetCurrency := ""
	valueInBRL := 0.0
	var err error

	fmt.Println("Bem vindo ao seu super 'Conversor de moedas!")
	fmt.Print("Insira a moeda para qual deseja converter: ")
	_, err = fmt.Scan(&targetCurrency)
	if err != nil {
		fmt.Println("Moeda não encontrada, confira a lista de compatibilidade: ", exchangeRates)
		return
	}

	fmt.Print("Insira o valor em reais que deseja converter: ")
	_, err = fmt.Scan(&valueInBRL)
	if err != nil {
		fmt.Println("Valor inválido!")
		return
	}

	targetCurrency = strings.ToUpper(targetCurrency)
	taxa, existe := exchangeRates[targetCurrency]
	if !existe {
		fmt.Println("Moeda não encontrada, confira a lista de compatibilidade: ", exchangeRates)
		return
	}

	if valueInBRL <= 0 {
		fmt.Printf("O valor deve ser maior que zero")
		return
	}

	convertedValue := valueInBRL * taxa	
	formattedValue := strconv.FormatFloat(convertedValue, 'f', 2, 64)
	formattedValue = strings.Replace(formattedValue, ".", ",", 1)

	fmt.Printf("Você está convertendo %.2f BRL para %s\n", valueInBRL, targetCurrency)
	fmt.Printf("O valor convertido é: %s %s\n", targetCurrency, formattedValue)
}