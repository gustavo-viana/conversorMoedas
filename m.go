package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Taxas struct {
	Base  string
	Date  string
	Rates map[string]float64
}

func main() {
	if len(os.Args) > 2 {

		file, err := os.ReadFile("rates.json")

		if err != nil {
			panic(err)
		}

		var js Taxas

		conversor := json.Unmarshal(file, &js)
		if conversor != nil {
			fmt.Println("Erro ao decodificar JSON")
			return
		}

		in, err := strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			fmt.Println("Impossivel converter, reveja se a entrada foi um numero")
			return
		}

		moeda := strings.ToUpper(os.Args[2])
		tax, find := js.Rates[moeda]

		if !find {
			fmt.Println("Moeda não encontrada")
			return
		}

		conv := calc(in, tax)
		saida := fmt.Sprintf("%.2f", conv)
		saida = strings.Replace(saida, ".", ",", 1)
		fmt.Printf("Valor convertido:%s %s \n", moeda, saida)
	} else {
		fmt.Println("È preciso o valor e a moeda para conversão")
	}
}

func calc(valor float64, tx float64) float64 {
	conv := valor * tx
	return conv
}
