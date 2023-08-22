package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"math/big"
)

func GenPubKeyHex(PrivateKey string) string {
	// Replace with your private key bytes
	privateKeyBytes, _ := hex.DecodeString(PrivateKey)

	// Convert the private key bytes to an ECDSA private key
	privKey := new(ecdsa.PrivateKey)
	privKey.Curve = elliptic.P256()
	privKey.D = new(big.Int).SetBytes(privateKeyBytes)
	privKey.PublicKey.Curve = privKey.Curve
	privKey.PublicKey.X, privKey.PublicKey.Y = privKey.Curve.ScalarBaseMult(privKey.D.Bytes())

	// Serialize the public key
	publicKeyBytes := elliptic.Marshal(elliptic.P256(), privKey.PublicKey.X, privKey.PublicKey.Y)
	publicKeyHex := hex.EncodeToString(publicKeyBytes)

	return publicKeyHex
}
