package cryptonat

import (
	utopiago "gopkg.in/sagleft/utopialib-go.v1"
)

//Handler is a handler for all requests for Crypton voucher management
type Handler struct {
	Client *utopiago.UtopiaClient
}
