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

	if *configPath == "" || *schemaPath == "" {
		fmt.Println("error: --config and --schema are required")
		os.Exit(2)
	}

	configFileBytes, err := os.ReadFile(*configPath)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
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
		os.Exit(2)
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
		val, ok := config[key]

		if !ok {
			fmt.Println("Missing key:", key)
			isValidConfig = false
			continue
		}

		typeInSchema := schemaBytes.Required[key]

		// compare type
		switch val.(type) {
		case string:
			if schemaBytes.Required[key] != "string" {
				fmt.Printf("type mismatch for %s, expected %s but received string\n", key, typeInSchema)
				isValidConfig = false
			}
		case float64:
			if schemaBytes.Required[key] != "number" {
				fmt.Printf("type mismatch for %s, expected %s but received float64\n", key, typeInSchema)
				isValidConfig = false
			}
		case bool:
			if schemaBytes.Required[key] != "bool" {
				fmt.Printf("type mismatch for %s, expected %s but received boolean\n", key, typeInSchema)
				isValidConfig = false
			}
		default:
			fmt.Println("unknown schema type:", typeInSchema)
			isValidConfig = false
		}
	}

	if isValidConfig {
		fmt.Println("config is valid")
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
