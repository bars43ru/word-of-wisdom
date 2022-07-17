package hashcash

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
)

func base64EncodeUInt(n uint) []byte {
	b := []byte(strconv.FormatUint(uint64(n), 10))
	return ([]byte)(base64.StdEncoding.EncodeToString(b))
}

func sha1Hash(s string) string {
	hash := sha256.New()
	_, err := io.WriteString(hash, s)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}
