module github.com/Du-bem/dominari_scriptum/main_node/server

go 1.24.1

replace github.com/Du-bem/dominari_scriptum/main_node/types => ../types

require (
	github.com/Du-bem/dominari_scriptum/main_node/types v0.0.0-00010101000000-000000000000
	github.com/bitcoin-sv/spv-wallet-go-client v1.0.0-beta.24
	github.com/bitcoin-sv/spv-wallet/models v1.0.0-beta.39
	github.com/btnguyen2k/consu/checksum v1.1.1
	github.com/gorilla/mux v1.8.1
)

require (
	github.com/bitcoin-sv/go-sdk v1.1.16 // indirect
	github.com/boombuler/barcode v1.0.1-0.20190219062509-6c824513bacc // indirect
	github.com/go-resty/resty/v2 v2.15.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pquerna/otp v1.4.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.33.0 // indirect
)
