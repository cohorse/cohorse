package models

import "crypto/rsa"

type School struct {
	Token string
	Classes []Class
	PublicKey rsa.PublicKey
	PrivateKey rsa.PrivateKey
	VotingKey string
}
