package token

import (
	"crypto"
	"time"
)


type Token interface {
	CreateTokenString(data any) (string , error)
	ExtractTokenData(tokenString string , data any) error 
}

type token struct {
	privateEd25519Key 	crypto.PrivateKey 
	publicEd25519Key 	crypto.PublicKey
	expiration 			time.Duration 
}