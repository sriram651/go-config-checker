package main

import (
	"encoding/json"
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

	var config map[string]interface{}

	unmarshalErr := json.Unmarshal(fileBytes, &config)

	if unmarshalErr != nil {
		fmt.Println("Error while unmarshalling the json:", unmarshalErr)
		os.Exit(2)
	}

	fmt.Println("The config is:", config)
}
