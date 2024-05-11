package core

import (
	"encoding/base64"
	"log"

	"github.com/kuingsmile/decodeGoogleOTP/utils"
)

func Flow(flags utils.Flags) error {

	dataURL, err := utils.DecodeQRImage(flags)
	if err != nil {
		log.Fatalf("Error reading QR code: %v", err)
		return err
	}

	data, err := base64.StdEncoding.DecodeString(dataURL)
	if err != nil {
		log.Fatalf("Error decoding base64: %v", err)
	}

	payload, err := utils.DecodeMessage(data)
	if err != nil {
		log.Fatalf("Error decoding message: %v", err)
	}
	if flags.Debug {
		log.Printf("Decoded message: %v", payload)
	}

	outputResult := utils.GenerateResult(payload)

	utils.PrintOTPCodes(outputResult, flags)
	utils.OutputToCSV(outputResult, flags)
	utils.OutputToJSON(outputResult, flags)
	utils.OutputToTxt(outputResult, flags)
	utils.OutputToURL(outputResult, flags)
	utils.OutputQRImage(outputResult, flags)
	return nil
}
