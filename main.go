package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	configPath := flag.String("config", "", "path to config file")
	// schemaPath := flag.String("schema", "", "path to schema file")
	flag.Parse()

	fileBytes, err := os.ReadFile(*configPath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Byte length:", len(fileBytes))
}
