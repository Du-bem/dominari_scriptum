module github.com/Du-bem/dominari_scriptum/main

go 1.24.1

replace (
	github.com/Du-bem/dominari_scriptum/main_node/db => ./db
	github.com/Du-bem/dominari_scriptum/main_node/server => ./server
	github.com/Du-bem/dominari_scriptum/main_node/types => ./types
)

require (
	github.com/Du-bem/dominari_scriptum/main_node/db v0.0.0-00010101000000-000000000000
	github.com/Du-bem/dominari_scriptum/main_node/server v0.0.0-00010101000000-000000000000
	github.com/Du-bem/dominari_scriptum/main_node/types v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
	github.com/mattn/go-sqlite3 v1.14.24
)

require (
	github.com/bitcoin-sv/go-sdk v1.1.16 // indirect
	github.com/bitcoin-sv/spv-wallet-go-client v1.0.0-beta.26 // indirect
	github.com/bitcoin-sv/spv-wallet/models v1.0.0-beta.39 // indirect
	github.com/boombuler/barcode v1.0.1-0.20190219062509-6c824513bacc // indirect
	github.com/btnguyen2k/consu/checksum v1.1.1 // indirect
	github.com/go-resty/resty/v2 v2.15.3 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pquerna/otp v1.4.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.33.0 // indirect
)
