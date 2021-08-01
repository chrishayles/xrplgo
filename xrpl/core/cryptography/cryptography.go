package cryptography

import (
	"errors"
	"fmt"
)

type CryptoImplementationType string

const (
	ED25519   CryptoImplementationType = "ED25519"
	SECP256K1 CryptoImplementationType = "SECP256K1"
)

//NewKeyPair -- Generates a new key pair of specified type.
func NewKeyPair(cit CryptoImplementationType, seed []byte, isValidator bool) (CryptoImplementation, error) {

	if cit == "SECP256K1" {
		s, err := newSecp256k1KeyPair(seed, isValidator)
		return s, err
	} else if cit == "ED25519" {
		e, err := newEd25519KeyPair(seed, isValidator)
		return e, err
	} else {
		return nil, errors.New(fmt.Sprintf("Invlaid CryptoImplementationType: %s", cit))
	}

}
