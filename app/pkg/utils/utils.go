package utils

import (
	"encoding/base64"
	"fmt"
)

func GetBase64Auth(account string, psw string) string {
	auth := fmt.Sprintf("%s:%s", account, psw)
	bytesString := []byte(auth)
	encodestr := base64.StdEncoding.EncodeToString(bytesString)
	token := fmt.Sprintf("Basic %s", encodestr)
	return token
}