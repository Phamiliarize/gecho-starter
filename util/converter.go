package util

import (
	b64 "encoding/base64"
	"strconv"
)

func B64FromUint32(v uint32) string {
	return b64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(int64(v), 10)))
}

func Uint32FromB64(v string) (uint32, error) {
	// Convert nextToken to requestCursor
	byteDecoded, err := b64.StdEncoding.DecodeString(v)
	if err != nil {
		return 0, err
	}

	decoded, err := strconv.ParseUint(string(byteDecoded), 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(decoded), nil
}
