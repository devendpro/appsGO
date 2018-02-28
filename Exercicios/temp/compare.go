package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	a := os.Args[1:2]

	fmt.Println(a)

	verdade := strings.Compare(a[0], "-a")

	fmt.Println(strings.Compare(a[0], "-a"))

	fmt.Println(verdade)
}
