package helpers

import (
	"crypto/md5"
	"encoding/hex"
)

func GetHash(image []byte) string {
	hash := md5.Sum(image)

	return hex.EncodeToString(hash[:])
}
