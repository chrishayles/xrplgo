package cryptography

import (
	"fmt"

	"golang.org/x/crypto/ed25519"
)

//Test - Placeholder function for library reference.
func Test() string {
	pub, pri, _ := ed25519.GenerateKey(nil)
	return fmt.Sprintf("Public: %s\n\nPrivate: %s\n", pub, pri)
}
