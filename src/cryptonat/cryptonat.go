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

//SetClient - connects another Utopia client to Handler
func (h *Handler) SetClient(client *utopiago.UtopiaClient) error {
	//TODO: check client connection
	h.Client = client
	return nil
}

//ActivateVoucher - an attempt to activate a voucher and get data about it
func (h *Handler) ActivateVoucher(voucherCode string) (ActivationData, error) {
	//TODO
	return ActivationData{}, nil
}

//CheckVoucherStatus - checks the voucher data, its activation status
func (h *Handler) CheckVoucherStatus(referenceNumber int) (VoucherData, error) {
	//TODO
	return VoucherData{}, nil
}

//GetVoucherAmount - asks for the amount of the voucher if it has already been activated
func (h *Handler) GetVoucherAmount(referenceNumber int) (float64, error) {
	//TOOD
	return 0, nil
}

//CreateVoucher - an attempt to create a voucher for a given amount
func (h *Handler) CreateVoucher(float64) error {
	//TODO
	return nil
}

//GetNetFee - asks for a commission to create a voucher with a specific amount
func (h *Handler) GetNetFee(amount float64) (float64, error) {
	//TODO: user data from client
	return amount * 0.0015, nil
}
