package crypto

import (
	"github.com/ProtonMail/gopenpgp/v2/helper"
)

func EncryptMessage(publicKey string, message string) (string, error) {
		armor, err := helper.EncryptMessageArmored(publicKey, message)
		if err != nil {
			return "", err
		}

		return string(armor), nil
}

func DecryptMessage(privateKey string, passphrase string, encryptedMessage string) (string, error) {
		decrypted, err := helper.DecryptMessageArmored(privateKey, []byte(passphrase), encryptedMessage)
		if err != nil {
			return "", err
		}

		return string(decrypted), nil
}
