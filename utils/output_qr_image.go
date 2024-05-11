package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/skip2/go-qrcode"
)

func OutputQRImage(outputResult []OutputResult, flags Flags) {
	if flags.QrDir == "" {
		return
	}

	// Create the directory if it doesn't exist
	if _, err := os.Stat(flags.QrDir); os.IsNotExist(err) {
		if err = os.Mkdir(flags.QrDir, os.ModePerm); err != nil {
			log.Fatalf("Failed to create directory: %s", err)
			return
		}
	}

	for i, otp := range outputResult {
		filename := fmt.Sprintf("%d-%s", i, otp.Name)
		if otp.Issuer != "" {
			filename = fmt.Sprintf("%s-%s.png", filename, otp.Issuer)
		} else {
			filename = fmt.Sprintf("%s.png", filename)
		}
		filename = strings.NewReplacer(":", "", "@", "", " ", "").Replace(filename)
		filename = filepath.Join(flags.QrDir, filename)

		if err := qrcode.WriteFile(otp.URL, qrcode.Medium, 256, filename); err != nil {
			log.Fatalf("Failed to write QR code to file: %s", err)
		} else {
			fmt.Printf("QR code %d written to %s\n", i, filename)
		}
	}
}
