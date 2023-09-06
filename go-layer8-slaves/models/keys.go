package models

type ECDHKeyExchangeRequest struct {
	PrivKeySlaveD []byte `json:"PrivKeySlaveD"`
}

type ECDHKeyExchangeOutput struct {
	PrivKeyServerD []byte `json:"PrivKeyServerD"`
}
