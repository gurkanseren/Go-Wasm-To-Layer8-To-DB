package models

import "math/big"

type ECDHKeyExchangeOutput struct {
	PubKeyServerX *big.Int `json:"pub_key_server_x"`
	PubKeyServerY *big.Int `json:"pub_key_server_y"`
}