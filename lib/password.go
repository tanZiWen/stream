package lib

import (
	//log "github.com/Sirupsen/logrus"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
)

const (
	ENCRYPT_ALGORITHM = "pbkdf2"
	ENCRYPT_HASHER = "sha256"
	ENCRYPT_TIMES = 16384
)


func EncryptPassword(password string, salts ...string) (encryptedPassword string, err error) {
	var salt string

	if len(salts) > 0 {
		salt = salts[0]

		if len(salt) != 6 {
			salt = RandomStr(6)
		}
	} else {
		salt = RandomStr(6)
	}

	dk, err := scrypt.Key([]byte(password), []byte(salt), ENCRYPT_TIMES, 8, 1, 32)

	if err != nil {
		return "", err
	}

	encryptedPassword = fmt.Sprintf("%s_%s_%s_%d_%s", ENCRYPT_ALGORITHM, ENCRYPT_HASHER, salt, ENCRYPT_TIMES, base64.StdEncoding.EncodeToString(dk))
	return encryptedPassword, nil
}
