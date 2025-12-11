package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: serve <dir>")
		os.Exit(1)
	}

	log.Println("Listening on http://localhost:8000")

	if err := http.ListenAndServe(
		"localhost:8000",
		http.FileServer(http.Dir(os.Args[1])),
	); err != nil {
		log.Fatalln(err)
	}
}
