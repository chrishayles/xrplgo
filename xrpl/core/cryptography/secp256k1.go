package cryptography

import (
	"fmt"

	ethSecp256k1 "github.com/ethereum/go-ethereum/crypto/secp256k1"
)

type Secp256k1 struct {
	seed      []byte
	private   []byte
	public    []byte
	validator bool
}

//Test -- Placeholder for package reference
func Test() {
	bc := ethSecp256k1.BitCurve{}
	fmt.Println(bc)
}

// IMPLEMENTS CRYPTOIMPLEMENTATION

//DeriveKeypair -- Derrive a new keypair from seed
func (s *Secp256k1) DeriveKeypair() error {
	/* 	var pub []byte
	   	var pri []byte
	   	var err error
	*/
	return nil
}

//Sign -- Sign message with private key.
func (s *Secp256k1) Sign(Message []byte) []byte {
	return nil
}

//IsValidMessage -- Validate message with public key and signature.
func (s *Secp256k1) IsValidMessage(Message, Signature []byte) bool {
	return false
}
