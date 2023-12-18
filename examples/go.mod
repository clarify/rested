module github.com/clarify/rested/examples

require (
	github.com/clarify/rested v0.1.1
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/justinas/alice v1.2.0
	github.com/rs/cors v1.8.0
	github.com/rs/zerolog v1.25.0
)

require (
	github.com/evanphx/json-patch v4.11.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rs/xid v1.3.0 // indirect
	golang.org/x/crypto v0.17.0 // indirect
)

replace github.com/clarify/rested => ./../

go 1.17
