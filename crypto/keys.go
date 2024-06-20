package crypto

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

type GPGKey struct {
    KeyID  string
    UserID string
}

func ListGPGKeys() ([]GPGKey, error) {
    cmd := exec.Command("gpg", "--list-keys", "--with-colons")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        return nil, err
    }

    var keys []GPGKey
    scanner := bufio.NewScanner(&out)
    var currentKey GPGKey
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "pub") {
            fields := strings.Split(line, ":")
            currentKey = GPGKey{
                KeyID: fields[4],
            }
        }
        if strings.HasPrefix(line, "uid") {
            fields := strings.Split(line, ":")
            currentKey.UserID = fields[9]
            keys = append(keys, currentKey)
        }
    }
    return keys, nil
}

func ListGPGPrivateKeys() ([]GPGKey, error) {
    cmd := exec.Command("gpg", "--list-secret-keys", "--with-colons")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        return nil, err
    }

    var keys []GPGKey
    scanner := bufio.NewScanner(&out)
    var currentKey GPGKey
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "sec") {
            fields := strings.Split(line, ":")
            currentKey = GPGKey{
                KeyID: fields[4],
            }
        }
        if strings.HasPrefix(line, "uid") {
            fields := strings.Split(line, ":")
            currentKey.UserID = fields[9]
            keys = append(keys, currentKey)
        }
    }
    return keys, nil
}

func ExportPublicKey(keyID string) (string, error) {
    cmd := exec.Command("gpg", "--armor", "--export", keyID)
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        return "", err
    }
    return out.String(), nil
}

func ExportPrivateKey(keyID string) (string, error) {
		cmd := exec.Command("gpg", "--armor", "--export-secret-key", keyID)
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			return "", err
		}
		return out.String(), nil
}
