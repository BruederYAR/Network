package input

import (
	"Network/pkg/protocol"
	"crypto/rsa"
	"io"
	"os"
)

func CreateDirectory(username string) error {
	path := pathToConnect + username + "/"

	if FileExist(path) {
		return nil
	}

	return os.MkdirAll(path, 0777)
}

func FileExist(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	return false
}

func SavePrivateKeyToFile(path string, key rsa.PrivateKey) error {
	data, err := protocol.PrivateKeyToBytes(key)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()
	file.Write(data)

	return nil
}

func LoadPrivateKeyOnFile(path string) (*rsa.PrivateKey, error) {
	file, err := os.Open(path)
	if err != nil {
		return &rsa.PrivateKey{}, err
	}
	defer file.Close()

	data := make([]byte, 9182)

	for {
		_, err := file.Read(data)
		if err == io.EOF {
			break
		}
	}

	key, err := protocol.BytesToPrivateKey(data)
	if err != nil {
		return &rsa.PrivateKey{}, err
	}

	return &key, nil
}
