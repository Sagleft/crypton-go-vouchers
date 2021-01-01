package cryptonat

import (
	"errors"

	//utopiago "gopkg.in/sagleft/utopialib-go.v2"
	utopiago "./utopiago" //for dev version
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
	ReferenceNumber int64   `json:"referenceNumber"`
	Amount          float64 `json:"amount"`
}

//Handler is a handler for all requests for Crypton voucher management
type Handler struct {
	Client *utopiago.UtopiaClient
}

//SetClient - connects another Utopia client to Handler
func (h *Handler) SetClient(client *utopiago.UtopiaClient) error {
	if !client.CheckClientConnection() {
		return errors.New("client disconnected")
	}
	h.Client = client
	return nil
}

//ActivateVoucher - an attempt to activate a voucher and get data about it
func (h *Handler) ActivateVoucher(voucherCode string) (ActivationData, error) {
	referenceNumber, err := h.Client.UseVoucher(voucherCode)
	if err != nil {
		return ActivationData{}, err
	}
	return h.CheckVoucherActivation(referenceNumber)
}

//CheckVoucherActivation - lite version of CheckVoucherStatus
func (h *Handler) CheckVoucherActivation(referenceNumber int64) (ActivationData, error) {
	data, err := h.CheckVoucherStatus(referenceNumber)
	if err != nil {
		return ActivationData{}, err
	}
	//TODO: add fields exists check
	return ActivationData{
		Status:          data.Status,
		ReferenceNumber: referenceNumber,
		Amount:          data.Amount,
	}, nil
}

//CheckVoucherStatus - checks the voucher data, its activation status
func (h *Handler) CheckVoucherStatus(referenceNumber int64) (VoucherData, error) {
	//TODO
	return VoucherData{}, nil
}

//GetVoucherAmount - asks for the amount of the voucher if it has already been activated
func (h *Handler) GetVoucherAmount(referenceNumber int) (float64, error) {
	//result := ActivationData{}
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
