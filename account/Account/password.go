package Account

import (
	"crypto/rand"
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

type password string

func generateSalt(len int) (bytes []byte, err error) {
	bytes = make([]byte, len)
	_, err = rand.Read(bytes)
	return
}

func (hashP *password) doHash(salt []byte) error {
	h := make([]byte, 64)
	c1 := sha3.NewCShake256([]byte(""), salt)
	if _, err := c1.Write([]byte(*hashP)); err != nil {
		return err
	}
	if _, err := c1.Read(h); err != nil {
		return err
	}
	*hashP = password(hex.EncodeToString(h))
	return nil
}
