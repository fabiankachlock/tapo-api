package klap

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strings"
)

type KLAPCipher struct {
	key []byte
	iv  []byte
	sig []byte
	seq int32
}

func NewCipher(localSeed, remoteSeed, userHash []byte) *KLAPCipher {
	localHash := append(append(localSeed, remoteSeed...), userHash...)
	iv, seq := ivDerive(localHash)

	return &KLAPCipher{
		key: keyDerive(localHash),
		iv:  iv,
		seq: seq,
		sig: sigDerive(localHash),
	}
}

func ivDerive(localHash []byte) (iv []byte, seq int32) {
	localHash = append([]byte("iv"), localHash...)
	shaSum := sha256.Sum256(localHash)
	iv = shaSum[:12]
	seq = int32(binary.BigEndian.Uint32(iv[12:]))
	return
}

func keyDerive(localHash []byte) []byte {
	localHash = append([]byte("lsk"), localHash...)
	hash := sha256.Sum256(localHash)
	return hash[:16]
}

func sigDerive(localHash []byte) []byte {
	localHash = append([]byte("ldk"), localHash...)
	hash := sha256.Sum256(localHash)
	return hash[:28]
}

func (c *KLAPCipher) Encrypt(data []byte) (payload []byte, seq int32, err error) {
	// increment seq
	c.seq += 1
	seq = c.seq

	// create cipher
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return []byte{}, 0, fmt.Errorf("failed to create aes cipher: %w", err)
	}

	// pad input data
	padSize := aes.BlockSize - (len(data) % aes.BlockSize)
	padding := strings.Repeat(string(rune(padSize)), padSize)
	paddedData := append(data, []byte(padding)...)

	// encrypt data
	cbc := cipher.NewCBCEncrypter(block, c.ivSeq(seq))
	cipherBytes := make([]byte, len(paddedData))
	cbc.CryptBlocks(cipherBytes, paddedData)

	// create signature
	intBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(intBytes, uint32(seq))
	signature := sha256.Sum256(
		append(append(c.sig, intBytes...), cipherBytes...),
	)

	return append(signature[:], cipherBytes...), seq, nil
}

func (c *KLAPCipher) Decrypt(data []byte) ([]byte, error) {
	// create cipher
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to create aes cipher: %w", err)
	}

	// decrypt data
	cbc := cipher.NewCBCDecrypter(block, c.ivSeq(c.seq))
	realBytes := make([]byte, len(data)-32)
	cbc.CryptBlocks(realBytes, data[32:])

	// remove padding
	padSize := int(realBytes[len(realBytes)-1])
	return realBytes[:(len(realBytes) - padSize)], nil
}

func (c *KLAPCipher) ivSeq(seq int32) []byte {
	intBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(intBytes, uint32(seq))
	return append(c.iv, intBytes...)
}
