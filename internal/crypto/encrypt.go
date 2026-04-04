package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

type EncryptedPayload struct {
	Algorithm  string `json:"algorithm"`
	Encoding   string `json:"encoding"`
	Nonce      string `json:"nonce"`
	Ciphertext string `json:"ciphertext"`
}

type Encryptor struct {
	gcm cipher.AEAD
}

func NewEncryptorFromBase64Key(b64Key string) (*Encryptor, error) {
	rawKey, err := base64.StdEncoding.DecodeString(b64Key)
	if err != nil {
		return nil, fmt.Errorf("failed to decode STATUS_ENCRYPTION_KEY base64: %w", err)
	}

	// AES-256 => 32 bytes
	if len(rawKey) != 32 {
		return nil, fmt.Errorf("invalid key length: got %d, expected 32 bytes (AES-256)", len(rawKey))
	}

	block, err := aes.NewCipher(rawKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	return &Encryptor{gcm: gcm}, nil
}

func (e *Encryptor) Encrypt(plaintext []byte) (*EncryptedPayload, error) {
	nonce := make([]byte, e.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}

	ciphertext := e.gcm.Seal(nil, nonce, plaintext, nil)

	return &EncryptedPayload{
		Algorithm:  "AES-256-GCM",
		Encoding:   "base64",
		Nonce:      base64.StdEncoding.EncodeToString(nonce),
		Ciphertext: base64.StdEncoding.EncodeToString(ciphertext),
	}, nil
}

