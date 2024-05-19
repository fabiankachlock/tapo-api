package passthrough

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
)

type PassthroughKeyPair struct {
	PrivateKey *rsa.PrivateKey
}

func NewPassthroughKeyPair() (*PassthroughKeyPair, error) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, err
	}

	return &PassthroughKeyPair{
		PrivateKey: key,
	}, nil
}

func (k *PassthroughKeyPair) GetPublicKey() []byte {
	pub := k.PrivateKey.Public()
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(pub.(*rsa.PublicKey)),
		},
	)
	return pubPEM
}

type PassthroughCipher struct {
	key []byte
	iv  []byte
}

func NewPassthroughCipher(key []byte, keyPair *PassthroughKeyPair) (*PassthroughCipher, error) {
	keyBytes := make([]byte, base64.StdEncoding.DecodedLen(len(key)))
	n, err := base64.StdEncoding.Decode(keyBytes, key)
	if err != nil {
		return nil, err
	}
	keyBytes = keyBytes[:n]

	buf, err := keyPair.PrivateKey.Decrypt(rand.Reader, key, rsa.DecryptPKCS1v15)
	if len(buf) != 32 {
		return nil, errors.New("expected 32 bytes")
	}

	return &PassthroughCipher{
		key: buf[0:16],
		iv:  buf[16:32],
	}, nil
}

func (c *PassthroughCipher) Encrypt(data []byte) ([]byte, error) {
	// create cipher
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return []byte{}, err
	}

	// pad input data
	fullSize := (len(data)/block.BlockSize() + 1) * block.BlockSize()
	paddedData := make([]byte, fullSize)
	copy(paddedData[0:len(data)], data)

	// encrypt data
	mode := cipher.NewCBCEncrypter(block, c.iv)
	cipherBytes := make([]byte, len(paddedData))
	mode.CryptBlocks(cipherBytes, paddedData)

	encodedCipher := base64.StdEncoding.EncodeToString(cipherBytes)
	return []byte(encodedCipher), nil
}

func (c *PassthroughCipher) Decrypt(data []byte) ([]byte, error) {
	decodedCipher := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(decodedCipher, data)
	if err != nil {
		return []byte{}, err
	}
	decodedCipher = decodedCipher[:n]

	// create cipher
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return []byte{}, err
	}

	// decrypt data
	mode := cipher.NewCBCDecrypter(block, c.iv)
	realBytes := make([]byte, len(decodedCipher))
	mode.CryptBlocks(realBytes, decodedCipher)

	return realBytes, nil
}

func (c *PassthroughCipher) DigestUsername(username string) []byte {
	h := sha1.New()
	h.Write([]byte(username))
	hash := h.Sum(nil)

	dst := make([]byte, hex.EncodedLen(len(hash)))
	hex.Encode(dst, hash)
	return dst
}
