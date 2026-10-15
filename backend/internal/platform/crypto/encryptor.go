package crypto

type Encryptable interface {
	SensitiveFields() []*string
	EntityName() string
}

type FieldEncryptor struct {
	crypto *Crypto
}

func NewFieldEncryptor(c *Crypto) *FieldEncryptor {
	return &FieldEncryptor{crypto: c}
}

// Secure cifra todos los campos sensibles
func (fe *FieldEncryptor) Secure(e Encryptable) error {
	aad := e.EntityName()
	for _, field := range e.SensitiveFields() {
		if field == nil || *field == "" || fe.crypto.IsEncrypted(*field) {
			continue
		}

		enc, err := fe.crypto.Encrypt(*field, aad)
		if err != nil {
			return err
		}
		*field = enc
	}
	return nil
}

// Reveal descifra todos los campos sensibles
func (fe *FieldEncryptor) Reveal(e Encryptable) error {
	aad := e.EntityName()
	for _, field := range e.SensitiveFields() {
		if field == nil || *field == "" || !fe.crypto.IsEncrypted(*field) {
			continue
		}

		dec, err := fe.crypto.Decrypt(*field, aad)
		if err != nil {
			return err
		}
		*field = dec
	}
	return nil
}

func (fe *FieldEncryptor) Hash(value string) string {
	return fe.crypto.Hash(value)
}

func (fe *FieldEncryptor) IsEncrypted(value string) bool {
	return fe.crypto.IsEncrypted(value)
}
