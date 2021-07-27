package cryptography

import (
	"bytes"
	"errors"

	stockEd25519 "crypto/ed25519"
)

//Creates new ed25519 key pair with validation
func newEd25519KeyPair(seed []byte, validator bool) (*Ed25519, error) {
	if validator {
		return nil, errors.New("ED25519 cannot be a validator.")
	}

	x := &Ed25519{
		seed:      seed,
		validator: false,
	}

	err := x.DeriveKeypair()

	return x, err
}

type Ed25519 struct {
	seed      []byte
	private   []byte
	public    []byte
	validator bool
}

// IMPLEMENTS CRYPTOIMPLEMENTATION

//DeriveKeypair -- Derrive a new keypair from seed
func (e *Ed25519) DeriveKeypair() error {
	var pub []byte
	var pri []byte
	var err error

	if e.validator {
		err = errors.New("ED25519 cannot be a validator.")
		return err
	}

	if e.seed != nil {
		r := bytes.NewReader(e.seed)
		pub, pri, err = stockEd25519.GenerateKey(r)
	} else {
		pub, pri, err = stockEd25519.GenerateKey(nil)
	}

	if err != nil {
		return err
	}

	e.public = pub
	e.private = pri

	return nil
}

//Sign -- Sign message with private key.
func (e *Ed25519) Sign(Message []byte) []byte {
	return stockEd25519.Sign(e.private, Message)
}

//IsValidMessage -- Validate message with public key and signature.
func (e *Ed25519) IsValidMessage(Message, Signature []byte) bool {
	return stockEd25519.Verify(e.public, Message, Signature)
}
