package mpesa

import (
	"encoding/base64"
	"fmt"
)

// GeneratePassword by base64 encoding BusinessShortcode, Passkey, and Timestamp
func GeneratePassword(shortCode, passkey, timestamp string) string {
	str := fmt.Sprintf("%s%s%s", shortCode, passkey, timestamp)
	return base64.StdEncoding.EncodeToString([]byte(str))
}
