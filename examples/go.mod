module github.com/searis/rest-layer/examples

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/justinas/alice v0.0.0-20171023064455-03f45bd4b7da
	github.com/rs/cors v1.6.0
	github.com/rs/zerolog v1.11.0
	github.com/searis/rest-layer v0.0.0-00010101000000-000000000000
	github.com/zenazn/goji v0.9.0 // indirect
)

require (
	github.com/evanphx/json-patch v4.1.0+incompatible // indirect
	github.com/rs/xid v1.3.0 // indirect
	golang.org/x/crypto v0.0.0-20181127143415-eb0de9b17e85 // indirect
)

replace github.com/searis/rest-layer => ./../

go 1.17
