package main

import (
	"log"
	"net/http"
)

func main() {
	// Informe o diretório onde estão os arquivos estáticos.
	fs := http.FileServer(http.Dir("/home/joe/myfilesStatic"))

	http.Handle("/", fs)

	log.Println("Listening...")

	// vai disponibilizar em http://localhost:8080
	http.ListenAndServe(":8080", nil)

}
