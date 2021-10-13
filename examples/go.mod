module github.com/searis/rested/examples

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/justinas/alice v1.2.0
	github.com/rs/cors v1.8.0
	github.com/rs/zerolog v1.25.0
	github.com/searis/rested v0.1.0-rc
)

require (
	github.com/evanphx/json-patch v4.11.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rs/xid v1.3.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
)

replace github.com/searis/rested => ./../

go 1.17
