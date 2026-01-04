package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gen <outDir>")
		os.Exit(1)
	}

	g := NewGenerator("assets")
	if err := g.Run(os.Args[1]); err != nil {
		fmt.Printf("generate: %v\n", err)
		os.Exit(1)
	}
}
