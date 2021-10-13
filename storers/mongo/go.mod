module github.com/searis/rested/storers/mongo

go 1.17

require (
	github.com/rs/cors v1.6.0
	github.com/searis/rested v0.0.0-00010101000000-000000000000
	gopkg.in/mgo.v2 v2.0.0-20160818020120-3f83fa500528
)

require (
	github.com/evanphx/json-patch v4.1.0+incompatible // indirect
	github.com/rs/xid v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20181127143415-eb0de9b17e85 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/searis/rested => ./../../
