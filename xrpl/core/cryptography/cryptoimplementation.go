package cryptography

type CryptoImplementation interface {
	DeriveKeypair(DecodedSeed string, IsValidator bool) ([]byte, []byte, error)
	Sign(Message []byte, PrivateKey []byte) ([]byte, error)
	IsValidMessage(Message, Signature, PublicKey []byte) bool
}
