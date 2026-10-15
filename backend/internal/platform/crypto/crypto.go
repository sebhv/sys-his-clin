package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"

	"io"
	"os"
	"strings"
)

const (
	ciphertextPrefix  = "enc"
	currentKeyVersion = "v1"
)

type Crypto struct {
	gcm     cipher.AEAD
	hmacKey []byte
}

func New() (*Crypto, error) {
	keyBytes := []byte(os.Getenv("ENCRYPTION_KEY"))
	if len(keyBytes) != 32 {
		return nil, errors.New("ENCRYPTION_KEY debe tener exactamente 32 caracteres")
	}

	block, err := aes.NewCipher([]byte(keyBytes))
	if err != nil {
		return nil, err
	}

	g, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	pepper := os.Getenv("HASH_PEPPER")
	if pepper == "" {
		return nil, errors.New("HASH_PEPPER no definido")
	}

	return &Crypto{gcm: g, hmacKey: []byte(pepper)}, nil
}

// Hash para busquedas y comparaciones
func (c *Crypto) Hash(value string) string {
	h := hmac.New(sha256.New, []byte(c.hmacKey))
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum(nil))
}

// IsEncrypted valida el formato: enc:<version>:<base64>
func (c *Crypto) IsEncrypted(v string) bool {

	parts := strings.SplitN(v, ":", 3)
	if len(parts) != 3 {
		return false
	}

	if parts[0] != ciphertextPrefix {
		return false
	}

	if parts[1] != currentKeyVersion {
		return false
	}

	//if !strings.HasPrefix(v, ciphertextPrefix) {
	//	return false
	//}
	_, err := base64.StdEncoding.DecodeString(parts[2])
	return err == nil
}

// Encrypt encripta (valga la redundancia)con AES-256-GCM. El AAD ata el ciphertext a su contexto
func (c *Crypto) Encrypt(plaintext, aad string) (string, error) {
	nonce := make([]byte, c.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// El nonce se antepone al ciphertext
	//ciphertext := c.gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	ciphertext := c.gcm.Seal(nonce, nonce, []byte(plaintext), []byte(aad))
	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	return fmt.Sprintf("%s:%s:%s", ciphertextPrefix, currentKeyVersion, encoded), nil
}

// Decrypt desencripta. Si el valor no está encriptado, lo devuelve tal cual
func (c *Crypto) Decrypt(encoded, aad string) (string, error) {
	if !c.IsEncrypted(encoded) {
		return encoded, nil
	}

	parts := strings.SplitN(encoded, ":", 3)
	data, err := base64.StdEncoding.DecodeString(parts[2])
	if err != nil {
		return "", err
	}

	nonceSize := c.gcm.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("ciphertext demasiado corto")
	}

	nonce := data[:nonceSize]
	ciphertext := data[nonceSize:]

	plaintext, err := c.gcm.Open(nil, nonce, ciphertext, []byte(aad))
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
