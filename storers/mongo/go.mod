module github.com/searis/rested/storers/mongo

go 1.17

require (
	github.com/rs/cors v1.6.0
	github.com/searis/rested v0.1.0-rc
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

require (
	github.com/evanphx/json-patch v4.1.0+incompatible // indirect
	github.com/rs/xid v1.3.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/searis/rested => ./../../
