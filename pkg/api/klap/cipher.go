package klap

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/binary"
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
	h := sha256.New()
	h.Write(localHash)
	hash := h.Sum(nil)
	iv = hash[0:12]
	seqBytes := hash[len(hash)-4:]
	seq = int32(binary.BigEndian.Uint32(seqBytes))
	return
}

func keyDerive(localHash []byte) []byte {
	localHash = append([]byte("lsk"), localHash...)
	h := sha256.New()
	h.Write(localHash)
	hash := h.Sum(nil)
	key := hash[0:16]
	return key
}

func sigDerive(localHash []byte) []byte {
	localHash = append([]byte("ldk"), localHash...)
	h := sha256.New()
	h.Write(localHash)
	hash := h.Sum(nil)
	sig := hash[0:28]
	return sig
}

func (c *KLAPCipher) Encrypt(data []byte) (payload []byte, seq int32, err error) {
	// increment seq
	c.seq += 1
	seq = c.seq + 1

	// create cipher
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return []byte{}, 0, err
	}

	// pad input data
	fullSize := (len(data)/block.BlockSize() + 1) * block.BlockSize()
	paddedData := make([]byte, fullSize)
	copy(paddedData[0:len(data)], data)

	// encrypt data
	mode := cipher.NewCBCEncrypter(block, c.ivSeq(seq))
	cipherBytes := make([]byte, len(paddedData))
	mode.CryptBlocks(cipherBytes, paddedData)

	// create signature
	intBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(intBytes, uint32(seq))
	signature := sha256.Sum256(
		append(append(c.sig, intBytes...), cipherBytes...),
	)

	return (append(signature[:], cipherBytes...)), seq, nil
}

func (c *KLAPCipher) Decrypt(seq int32, data []byte) ([]byte, error) {
	// create cipher
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return []byte{}, err
	}

	// decrypt data
	mode := cipher.NewCBCDecrypter(block, c.ivSeq(seq))
	realBytes := make([]byte, len(data)-32)
	mode.CryptBlocks(realBytes, data[32:])

	// remove padding
	padSize := int(realBytes[len(realBytes)-1])
	return realBytes[:(len(realBytes) - padSize)], nil
}

func (c *KLAPCipher) ivSeq(seq int32) []byte {
	intBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(intBytes, uint32(seq))
	return append(c.iv, intBytes...)
}
