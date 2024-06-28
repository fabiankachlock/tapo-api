package tapoutil

import "encoding/base64"

func GetNickname(nickname string) string {
	decodedBytes, err := base64.StdEncoding.DecodeString(nickname)
	if err != nil {
		return nickname
	}
	return string(decodedBytes)
}
