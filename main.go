package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Schema struct {
	Required map[string]string
}

func main() {
	configPath := flag.String("config", "", "path to config file")
	schemaPath := flag.String("schema", "", "path to schema file")
	flag.Parse()

	configFileBytes, err := os.ReadFile(*configPath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var config map[string]interface{}

	unmarshalErr := json.Unmarshal(configFileBytes, &config)

	if unmarshalErr != nil {
		fmt.Println("Error while unmarshalling the json:", unmarshalErr)
		os.Exit(2)
	}

	schemaFileBytes, schemaFileErr := os.ReadFile(*schemaPath)

	if schemaFileErr != nil {
		fmt.Println("Error while reading schema.json:", schemaFileErr)
		os.Exit(1)
	}

	var schemaBytes Schema

	schemaUnmarshalErr := json.Unmarshal(schemaFileBytes, &schemaBytes)

	if schemaUnmarshalErr != nil {
		fmt.Println("Error unmarshalling schema.json:", schemaUnmarshalErr)
		os.Exit(2)
	}

	fmt.Println("Required keys:", len(schemaBytes.Required))

	isValidConfig := true
	for key := range schemaBytes.Required {
		_, ok := config[key]

		if ok != true {
			fmt.Println("Missing key:", key)
			isValidConfig = false
		}
	}

	if isValidConfig {
		fmt.Println("All required keys are present!")
	}
}
