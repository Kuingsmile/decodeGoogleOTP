package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func OutputToTxt(outputResult []OutputResult, flags Flags) {
	if flags.Text == "" {
		return
	}

	filename := flags.Text
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
		return
	}
	defer file.Close()

	// Create a writer from the file
	writer := io.MultiWriter(file)

	fmt.Fprintln(writer, "-----------------------------------")
	for _, otp := range outputResult {
		var builder strings.Builder
		builder.WriteString(fmt.Sprintf("issuer: %s\n", otp.Issuer))
		builder.WriteString(fmt.Sprintf("name: %s\n", otp.Name))
		builder.WriteString(fmt.Sprintf("secret: %s\n", otp.Secret))
		builder.WriteString(fmt.Sprintf("type: %s\n", otp.Type))
		builder.WriteString(fmt.Sprintf("counter: %d\n", otp.Counter))
		builder.WriteString(fmt.Sprintf("url: %s\n", otp.URL))
		builder.WriteString("-----------------------------------\n")
		fmt.Fprint(writer, builder.String())
	}
	fmt.Printf("Text results written to %s\n", filename)
}
