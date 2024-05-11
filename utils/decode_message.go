package utils

import (
	"github.com/kuingsmile/decodeGoogleOTP/migrationpayload"
	"google.golang.org/protobuf/proto"
)

func DecodeMessage(buffer []byte) (*migrationpayload.MigrationPayload, error) {
	payload := &migrationpayload.MigrationPayload{}
	if err := proto.Unmarshal(buffer, payload); err != nil {
		return nil, err
	}
	return payload, nil
}
