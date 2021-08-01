package cryptography

type CryptoImplementation interface {
	DeriveKeypair() error
	Sign(Message []byte) []byte
	IsValidMessage(Message, Signature []byte) bool
}
