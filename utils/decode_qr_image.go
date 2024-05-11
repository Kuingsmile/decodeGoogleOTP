package utils

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/url"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func DecodeQRImage(flags Flags) (string, error) {

	filePath := flags.Input
	if filePath == "" {
		log.Fatalf("No input file provided")
	}

	dataURL, err := readQRCodeImage(filePath)
	if err != nil {
		return "", err
	}

	fmt.Println("Decoded QR code successfully")
	if flags.Debug {
		fmt.Printf("Decoded data: %s\n", dataURL)
	}
	return dataURL, nil
}

func readQRCodeImage(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	fmt.Println("File opened successfully")
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return "", err
	}

	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		return "", err
	}

	if result.String() == "" {
		return "", fmt.Errorf("no qr code found")
	}

	length := len("otpauth-migration://offline?data=")

	if len(result.String()) < length {
		return "", fmt.Errorf("invalid QR code")
	}

	return url.QueryUnescape(result.String()[length:])

}
