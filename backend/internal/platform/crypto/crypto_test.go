package crypto

import (
	"os"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	os.Setenv("ENCRYPTION_KEY", "12345678901234567890123456789012")
	os.Setenv("HASH_PEPPER", "mi-pepper-secreto")

	c, err := New()
	if err != nil {
		t.Fatal(err)
	}

	plain := "juan@example.com"
	aad := "paciente"

	enc, err := c.Encrypt(plain, aad)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("encrypted: %s", enc)

	if !c.IsEncrypted(enc) {
		t.Fatal("IsEncrypted devolvió false")
	}

	dec, err := c.Decrypt(enc, aad)
	if err != nil {
		t.Fatal(err)
	}
	if dec != plain {
		t.Fatalf("esperado %q, obtenido %q", plain, dec)
	}

	// Verificar que el AAD incorrecto falla
	if _, err := c.Decrypt(enc, "personal"); err == nil {
		t.Fatal("esperaba error con AAD incorrecto")
	}
}
