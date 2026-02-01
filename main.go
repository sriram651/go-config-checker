package main

import (
	"flag"
	"fmt"
)

func main() {
	configPath := flag.String("config", "", "path to config file")
	schemaPath := flag.String("schema", "", "path to schema file")
	flag.Parse()

	fmt.Println("config:", *configPath)
	fmt.Println("schema:", *schemaPath)
}
