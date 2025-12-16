package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	tx := 2.5
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite o valor: ")
	input, _ := reader.ReadString('\n')
	cleanerValue, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Valor invalido")
		return
	}
	convertValue := calc(cleanerValue, tx)
	fmt.Printf("Valor convetido:$%.2f\n", convertValue)
}

func calc(valor float64, tx float64) float64 {
	imposto := valor * (tx / 100)
	cabio := 5.50
	conv := (valor - imposto) / cabio
	return conv
}
