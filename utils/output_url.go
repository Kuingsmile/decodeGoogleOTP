package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func OutputToURL(outputResult []OutputResult, flags Flags) {
	if flags.Url == "" {
		return
	}

	filename := flags.Url
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
		return
	}
	defer file.Close()

	// Create a writer from the file
	writer := io.MultiWriter(file)

	for _, otp := range outputResult {
		var builder strings.Builder
		builder.WriteString(fmt.Sprintf("url: %s\n", otp.URL))
		fmt.Fprint(writer, builder.String())
	}
	fmt.Printf("URL results written to %s\n", filename)
}
