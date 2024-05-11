package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func OutputToCSV(outputResult []OutputResult, flags Flags) {
	if flags.Csv == "" {
		return
	}
	filename := flags.Csv
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Issuer", "Name", "Secret", "Type", "Counter", "URL"})

	for _, otp := range outputResult {
		writer.Write([]string{otp.Issuer, otp.Name, otp.Secret, otp.Type, fmt.Sprintf("%d", otp.Counter), otp.URL})
	}

	err = writer.Error()
	if err != nil {
		log.Fatalf("Failed to write to file: %s", err)
	}
	fmt.Printf("CSV results written to %s\n", filename)
}
