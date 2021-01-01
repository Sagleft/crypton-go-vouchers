# crypton-go-vouchers
Crypton Donation Library used to work with vouchers written in Go

## Installation

```bash
go get github.com/Sagleft/crypton-go-vouchers/cryptonat
```

then

```go
import "github.com/Sagleft/crypton-go-vouchers/cryptonat"
```

Usage
-------

```go
//TODO
```

Then try activating the voucher:

```go
//TODO
```

response example:

```json
{
	"status": "pending",
	"referenceNumber": "367404A95932624C284B16AF1C1EDF1BB0F9CDCA1CC5136B167378BBF933FAD8",
	"amount": 0
}
```

check voucher status by reference number:

```go
//TODO
```

response example:

```json
{
	"status": "done",
	"created": "2020-01-14T13:18:21.232",
	"amount": 2,
	"comments": "",
	"direction": 1,
	"trid": "0ZWTT62Z4DO51"
}
```

![scheme](https://github.com/Sagleft/crypton-go-vouchers/raw/main/assets/voucher_activation.png)

# How can this be used?
* creating a web service that processes client requests;
* creation of a payment service;
* development of a bot for the channel.

License
-------

crypton-go-vouchers is licensed under [The MIT License](LICENSE).
