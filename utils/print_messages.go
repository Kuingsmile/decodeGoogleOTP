package utils

import (
	"fmt"
	"strings"

	"github.com/mdp/qrterminal/v3"
)

func PrintOTPCodes(outputResult []OutputResult, flags Flags) {
	if flags.Silent {
		return
	}
	fmt.Println("Results:")
	fmt.Println("-----------------------------------")
	for _, otp := range outputResult {

		var builder strings.Builder
		builder.WriteString(fmt.Sprintf("issuer: %s\n", otp.Issuer))
		builder.WriteString(fmt.Sprintf("name: %s\n", otp.Name))
		builder.WriteString(fmt.Sprintf("secret: %s\n", otp.Secret))
		builder.WriteString(fmt.Sprintf("type: %s\n", otp.Type))
		builder.WriteString(fmt.Sprintf("counter: %d\n", otp.Counter))
		builder.WriteString(fmt.Sprintf("url: %s\n\n", otp.URL))

		if flags.PrintQR {

			config := qrterminal.Config{
				Level:     qrterminal.M,
				Writer:    &builder,
				BlackChar: qrterminal.WHITE,
				WhiteChar: qrterminal.BLACK,
				QuietZone: 3,
			}
			qrterminal.GenerateWithConfig(otp.URL, config)

			builder.WriteString("-----------------------------------\n")
		}
		fmt.Print(builder.String())
	}
}
