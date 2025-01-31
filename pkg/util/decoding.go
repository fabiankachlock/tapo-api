package tapoutil

import "encoding/base64"

func decodeBase64(value string) string {
	decodedBytes, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return value
	}
	return string(decodedBytes)
}

func GetNickname(nickname string) string {
	return decodeBase64(nickname)
}

func GetSSID(nickname string) string {
	return decodeBase64(nickname)
}
