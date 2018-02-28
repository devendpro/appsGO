package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	ascDesc := os.Args[1:2]
	//ascendente := ascDesc[0] == "-a"
	ascendente := strings.Compare(ascDesc[0],"-a") == 0

	fmt.Printf("O primeiro caracter é:\"%s\"\n", ascDesc[0])

	if (ascDesc[0] != "-a") && (ascDesc[0] != "-b") {
		fmt.Println("O primeiro parâmetro da função deve ser o caracter -a ou -b para indicar a ordenação dos números.")
		os.Exit(1)

	} else {

		entrada := os.Args[2:]
		numeros := make([]int, len(entrada))

		for i, n := range entrada {
			numero, err := strconv.Atoi(n)
			if err != nil {
				fmt.Printf("%s não é um número válido", n)
				os.Exit(1)
			}
			numeros[i] = numero
		}

		fmt.Println(quicksort(numeros, ascendente))

	}

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

func quicksort(numeros []int, ascendente bool) []int {
	if len(numeros) <= 1 {
		return numeros
	}

	indicePivo := len(numeros) / 2
	pivo := numeros[indicePivo]
	numeros = append(numeros[:indicePivo], numeros[indicePivo+1:]...)

	menores, maiores := particionar(numeros, pivo)

	return append(append(quicksort(maiores, ascendente), pivo), quicksort(menores, ascendente)...)

}
