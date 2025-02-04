package keystore

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
)

type KeyStore struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

func NewKeyStore() *KeyStore {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	publicKey := &privateKey.PublicKey
	return &KeyStore{privateKey, publicKey}
}

func (ks *KeyStore) GetAddress() string {
	pubKeyBytes := elliptic.Marshal(ks.PrivateKey.Curve, ks.PublicKey.X, ks.PublicKey.Y)
	return hex.EncodeToString(pubKeyBytes)
}