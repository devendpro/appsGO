package main

import (
	"fmt"
	"net/http"
)

// PORT é a porta
const (
	PORT = ":8080"
)

func main() {
	fmt.Println("Executando...")
	http.ListenAndServe(PORT, http.FileServer(http.Dir(".")))

}
