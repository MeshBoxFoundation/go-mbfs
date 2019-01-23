// Package crypto implements various cryptographic utilities used by ipfs.
// This includes a Public and Private key interface and an RSA key implementation
// that satisfies it.
package crypto

import (
	"bytes"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
	"io"

	pb "mbfs/go-mbfs/gx/QmNiJiXwWE3kRhZrC5ej3kSjWHm337pYfhjLGSCDNKJP2s/go-libp2p-crypto/pb"

	sha256 "mbfs/go-mbfs/gx/QmcTzQXRcU2vf8yX5EEboz1BSvWC7wWmeYAKVQmhp8WZYU/sha256-simd"
	"mbfs/go-mbfs/gx/QmdxUuburamoF6zF9qjeQC4WYcWGbWuRmdLacMEsW8ioD8/gogo-protobuf/proto"
)

const (
	// RSA is an enum for the supported RSA key type
	RSA = iota
	// Ed25519 is an enum for the supported Ed25519 key type
	Ed25519
	// Secp256k1 is an enum for the supported Secp256k1 key type
	Secp256k1
	// ECDSA is an enum for the supported ECDSA key type
	ECDSA
)

var (
	// ErrBadKeyType is returned when a key is not supported
	ErrBadKeyType = errors.New("invalid or unsupported key type")
	// KeyTypes is a list of supported keys
	KeyTypes = []int{
		RSA,
		Ed25519,
		Secp256k1,
		ECDSA,
	}
)

// PubKeyUnmarshaller is a func that creates a PubKey from a given slice of bytes
type PubKeyUnmarshaller func(data []byte) (PubKey, error)

// PrivKeyUnmarshaller is a func that creates a PrivKey from a given slice of bytes
type PrivKeyUnmarshaller func(data []byte) (PrivKey, error)

// PubKeyUnmarshallers is a map of unmarshallers by key type
var PubKeyUnmarshallers = map[pb.KeyType]PubKeyUnmarshaller{
	pb.KeyType_RSA:       UnmarshalRsaPublicKey,
	pb.KeyType_Ed25519:   UnmarshalEd25519PublicKey,
	pb.KeyType_Secp256k1: UnmarshalSecp256k1PublicKey,
	pb.KeyType_ECDSA:     UnmarshalECDSAPublicKey,
}

// PrivKeyUnmarshallers is a map of unmarshallers by key type
var PrivKeyUnmarshallers = map[pb.KeyType]PrivKeyUnmarshaller{
	pb.KeyType_RSA:       UnmarshalRsaPrivateKey,
	pb.KeyType_Ed25519:   UnmarshalEd25519PrivateKey,
	pb.KeyType_Secp256k1: UnmarshalSecp256k1PrivateKey,
	pb.KeyType_ECDSA:     UnmarshalECDSAPrivateKey,
}

// Key represents a crypto key that can be compared to another key
type Key interface {
	// Bytes returns a serialized, storeable representation of this key
	// DEPRECATED in favor of Marshal / Unmarshal
	Bytes() ([]byte, error)

	// Equals checks whether two PubKeys are the same
	Equals(Key) bool

	// Raw returns the raw bytes of the key (not wrapped in the
	// libp2p-crypto protobuf).
	//
	// This function is the inverse of {Priv,Pub}KeyUnmarshaler.
	Raw() ([]byte, error)

	// Type returns the protobof key type.
	Type() pb.KeyType
}

// PrivKey represents a private key that can be used to generate a public key,
// sign data, and decrypt data that was encrypted with a public key
type PrivKey interface {
	Key

	// Cryptographically sign the given bytes
	Sign([]byte) ([]byte, error)

	// Return a public key paired with this private key
	GetPublic() PubKey
}

// PubKey is a public key
type PubKey interface {
	Key

	// Verify that 'sig' is the signed hash of 'data'
	Verify(data []byte, sig []byte) (bool, error)
}

// GenSharedKey generates the shared key from a given private key
type GenSharedKey func([]byte) ([]byte, error)

// GenerateKeyPair generates a private and public key
func GenerateKeyPair(typ, bits int) (PrivKey, PubKey, error) {
	return GenerateKeyPairWithReader(typ, bits, rand.Reader)
}

// GenerateKeyPairWithReader returns a keypair of the given type and bitsize
func GenerateKeyPairWithReader(typ, bits int, src io.Reader) (PrivKey, PubKey, error) {
	switch typ {
	case RSA:
		return GenerateRSAKeyPair(bits, src)
	case Ed25519:
		return GenerateEd25519Key(src)
	case Secp256k1:
		return GenerateSecp256k1Key(src)
	case ECDSA:
		return GenerateECDSAKeyPair(src)
	default:
		return nil, nil, ErrBadKeyType
	}
}

// GenerateEKeyPair returns an ephemeral public key and returns a function that will compute
// the shared secret key.  Used in the identify module.
//
// Focuses only on ECDH now, but can be made more general in the future.
func GenerateEKeyPair(curveName string) ([]byte, GenSharedKey, error) {
	var curve elliptic.Curve

	switch curveName {
	case "P-256":
		curve = elliptic.P256()
	case "P-384":
		curve = elliptic.P384()
	case "P-521":
		curve = elliptic.P521()
	}

	priv, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	pubKey := elliptic.Marshal(curve, x, y)

	done := func(theirPub []byte) ([]byte, error) {
		// Verify and unpack node's public key.
		x, y := elliptic.Unmarshal(curve, theirPub)
		if x == nil {
			return nil, fmt.Errorf("malformed public key: %d %v", len(theirPub), theirPub)
		}

		if !curve.IsOnCurve(x, y) {
			return nil, errors.New("invalid public key")
		}

		// Generate shared secret.
		secret, _ := curve.ScalarMult(x, y, priv)

		return secret.Bytes(), nil
	}

	return pubKey, done, nil
}

// StretchedKeys ...
type StretchedKeys struct {
	IV        []byte
	MacKey    []byte
	CipherKey []byte
}

// KeyStretcher returns a set of keys for each party by stretching the shared key.
// (myIV, theirIV, myCipherKey, theirCipherKey, myMACKey, theirMACKey)
func KeyStretcher(cipherType string, hashType string, secret []byte) (StretchedKeys, StretchedKeys) {
	var cipherKeySize int
	var ivSize int
	switch cipherType {
	case "AES-128":
		ivSize = 16
		cipherKeySize = 16
	case "AES-256":
		ivSize = 16
		cipherKeySize = 32
	case "Blowfish":
		ivSize = 8
		// Note: 24 arbitrarily selected, needs more thought
		cipherKeySize = 32
	}

	hmacKeySize := 20

	seed := []byte("key expansion")

	result := make([]byte, 2*(ivSize+cipherKeySize+hmacKeySize))

	var h func() hash.Hash

	switch hashType {
	case "SHA1":
		h = sha1.New
	case "SHA256":
		h = sha256.New
	case "SHA512":
		h = sha512.New
	default:
		panic("Unrecognized hash function, programmer error?")
	}

	m := hmac.New(h, secret)
	// note: guaranteed to never return an error
	m.Write(seed)

	a := m.Sum(nil)

	j := 0
	for j < len(result) {
		m.Reset()

		// note: guaranteed to never return an error.
		m.Write(a)
		m.Write(seed)

		b := m.Sum(nil)

		todo := len(b)

		if j+todo > len(result) {
			todo = len(result) - j
		}

		copy(result[j:j+todo], b)

		j += todo

		m.Reset()

		// note: guaranteed to never return an error.
		m.Write(a)

		a = m.Sum(nil)
	}

	half := len(result) / 2
	r1 := result[:half]
	r2 := result[half:]

	var k1 StretchedKeys
	var k2 StretchedKeys

	k1.IV = r1[0:ivSize]
	k1.CipherKey = r1[ivSize : ivSize+cipherKeySize]
	k1.MacKey = r1[ivSize+cipherKeySize:]

	k2.IV = r2[0:ivSize]
	k2.CipherKey = r2[ivSize : ivSize+cipherKeySize]
	k2.MacKey = r2[ivSize+cipherKeySize:]

	return k1, k2
}

// UnmarshalPublicKey converts a protobuf serialized public key into its
// representative object
func UnmarshalPublicKey(data []byte) (PubKey, error) {
	pmes := new(pb.PublicKey)
	err := proto.Unmarshal(data, pmes)
	if err != nil {
		return nil, err
	}

	um, ok := PubKeyUnmarshallers[pmes.GetType()]
	if !ok {
		return nil, ErrBadKeyType
	}

	return um(pmes.GetData())
}

// MarshalPublicKey converts a public key object into a protobuf serialized
// public key
func MarshalPublicKey(k PubKey) ([]byte, error) {
	pbmes := new(pb.PublicKey)
	pbmes.Type = k.Type()
	data, err := k.Raw()
	if err != nil {
		return nil, err
	}
	pbmes.Data = data

	return proto.Marshal(pbmes)
}

// UnmarshalPrivateKey converts a protobuf serialized private key into its
// representative object
func UnmarshalPrivateKey(data []byte) (PrivKey, error) {
	pmes := new(pb.PrivateKey)
	err := proto.Unmarshal(data, pmes)
	if err != nil {
		return nil, err
	}

	um, ok := PrivKeyUnmarshallers[pmes.GetType()]
	if !ok {
		return nil, ErrBadKeyType
	}

	return um(pmes.GetData())
}

// MarshalPrivateKey converts a key object into its protobuf serialized form.
func MarshalPrivateKey(k PrivKey) ([]byte, error) {
	pbmes := new(pb.PrivateKey)
	pbmes.Type = k.Type()
	data, err := k.Raw()
	if err != nil {
		return nil, err
	}

	pbmes.Data = data
	return proto.Marshal(pbmes)
}

// ConfigDecodeKey decodes from b64 (for config file), and unmarshals.
func ConfigDecodeKey(b string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(b)
}

// ConfigEncodeKey encodes to b64 (for config file), and marshals.
func ConfigEncodeKey(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// KeyEqual checks whether two
func KeyEqual(k1, k2 Key) bool {
	if k1 == k2 {
		return true
	}

	b1, err1 := k1.Bytes()
	b2, err2 := k2.Bytes()
	return bytes.Equal(b1, b2) && err1 == err2
}
