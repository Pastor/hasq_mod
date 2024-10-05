package hashq_mod

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"math/big"
)

type DeviceCrypto struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	Original   *ecdsa.PrivateKey
}

type SignatureKey struct {
	R, S *big.Int
}

func (dk DeviceCrypto) Public() crypto.PublicKey {
	return dk.Original.Public()
}

func (dk DeviceCrypto) Sign(content []byte) string {
	digest := dk.Digest(content)
	r, s, err := ecdsa.Sign(rand.Reader, dk.Original, digest)
	if err != nil {
		panic(err)
	}
	var sig = SignatureKey{S: s, R: r}
	var signature, asnErr = asn1.Marshal(sig)
	if asnErr != nil {
		panic(asnErr)
	}
	return EncodeToString(signature)
}

func (dk DeviceCrypto) Verify(pub crypto.PublicKey, content []byte, signature []byte) bool {
	var sig SignatureKey
	digest := dk.Digest(content)
	_, _ = asn1.Unmarshal(signature, &sig)
	if pub == nil {
		pub = dk.Public()
	}
	return ecdsa.Verify(pub.(*ecdsa.PublicKey), digest, sig.R, sig.S)
}

func (dk DeviceCrypto) Digest(data []byte) []byte {
	bytes := sha256.Sum256(data)
	return bytes[:]
}

func GenerateKey() DeviceCrypto {
	key, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		panic(err)
	}
	privateKey, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		panic(err)
	}

	publicKey, err := x509.MarshalPKIXPublicKey(key.Public())
	if err != nil {
		panic(err)
	}
	return DeviceCrypto{
		PrivateKey: EncodeToString(privateKey),
		PublicKey:  EncodeToString(publicKey),
		Original:   key,
	}
}
