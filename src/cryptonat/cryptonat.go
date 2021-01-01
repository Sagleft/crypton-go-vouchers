package cryptonat

import (
	utopiago "gopkg.in/sagleft/utopialib-go.v1"
)

//VoucherData contains voucher data ^ↀᴥↀ^
type VoucherData struct {
	Status           string  `json:"status"`
	CreatedTimestamp string  `json:"created"`
	Amount           float64 `json:"amount"`
	Comments         string  `json:"comments"`
	Direction        int     `json:"direction"`
	TransactionID    string  `json:"trid"`
}

//ActivationData contains voucher activation data (=ↀωↀ=)
type ActivationData struct {
	Status          string  `json:"status"`
	ReferenceNumber int     `json:"referenceNumber"`
	Amount          float64 `json:"amount"`
}

//Handler is a handler for all requests for Crypton voucher management
type Handler struct {
	Client *utopiago.UtopiaClient
}
