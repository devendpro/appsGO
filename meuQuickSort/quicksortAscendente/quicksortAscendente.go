package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	entrada := os.Args[1:]
	numeros := make([]int, len(entrada))

	for i, n := range entrada {
		numero, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s não é um número válido", n)
			os.Exit(1)
		}
		numeros[i] = numero
	}

	fmt.Println(quicksort(numeros))

}

func particionar(numeros []int, pivo int) ([]int, []int) {
	menores := []int{}
	maiores := []int{}
	for _, n := range numeros {
		if n <= pivo {
			menores = append(menores, n)
		} else {
			maiores = append(maiores, n)
		}
	}

	return menores, maiores
}

func quicksort(numeros []int) []int {
	if len(numeros) <= 1 {
		return numeros
	}

	indicePivo := len(numeros) / 2
	pivo := numeros[indicePivo]
	numeros = append(numeros[:indicePivo], numeros[indicePivo+1:]...)

	menores, maiores := particionar(numeros, pivo)

	return append(append(quicksort(maiores), pivo), quicksort(menores)...)

}
