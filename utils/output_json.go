package utils

import (
	"encoding/json"
	"log"
	"os"
)

func OutputToJSON(outputResult []OutputResult, flags Flags) {
	if flags.Json == "" {
		return
	}
	filename := flags.Json
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(outputResult); err != nil {
		log.Fatalf("Failed to write to file: %s", err)
	} else {
		log.Printf("JSON results written to %s\n", filename)
	}
}
