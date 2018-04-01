package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "William de Jesus Pinheiro Martins"
	scanner := bufio.NewScanner(strings.NewReader(s))

	scanner.Split(bufio.ScanRunes)

	palavrinhas := make([]rune,1)
	i := 0

	for scanner.Scan() {
		palavrinhas = append(palavrinhas, scanner.Text())
		i++
	}


	fmt.Println(palavrinhas)
	fmt.Println(palavrinhas[1],palavrinhas[5])
}

