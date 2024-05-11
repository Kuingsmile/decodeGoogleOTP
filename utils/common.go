package utils

import (
	"encoding/base32"
	"net/url"
	"strings"

	"github.com/kuingsmile/decodeGoogleOTP/migrationpayload"
)

func getOTPTypeString(otpType string) string {
	switch otpType {
	case "OTP_HOTP":
		return "hotp"
	case "OTP_TOTP":
		return "totp"
	default:
		return "totp"
	}
}

func GenerateURL(issuer, name, secret, othType string) string {
	var sb strings.Builder
	sb.WriteString("otpauth://")
	sb.WriteString(getOTPTypeString(othType))
	sb.WriteString("/")
	sb.WriteString(url.QueryEscape(name))
	sb.WriteString("?secret=")
	sb.WriteString(secret)

	if issuer != "" {
		sb.WriteString("&issuer=")
		sb.WriteString(url.QueryEscape(issuer))
	}

	return sb.String()
}

func GenerateResult(payload *migrationpayload.MigrationPayload) []OutputResult {
	var results []OutputResult
	for _, otp := range payload.OtpParameters {
		secret := base32.StdEncoding.EncodeToString(otp.Secret)
		secret = strings.TrimRight(secret, "=")
		otpType := getOTPTypeString(string(otp.Type))
		url := GenerateURL(otp.Issuer, otp.Name, secret, otpType)
		results = append(results, OutputResult{
			Issuer:  otp.Issuer,
			Name:    otp.Name,
			Secret:  secret,
			Type:    otpType,
			Counter: otp.Counter,
			URL:     url,
		})
	}
	return results
}
