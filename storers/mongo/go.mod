module github.com/rs/rest-layer-mongo

go 1.15

require (
	github.com/kr/pretty v0.1.0 // indirect
	github.com/rs/cors v1.6.0
	github.com/searis/rested v0.0.0-20210930125824-35d9e099e82d
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/mgo.v2 v2.0.0-20160818020120-3f83fa500528
	gopkg.in/yaml.v2 v2.2.2 // indirect
)

replace github.com/searis/rested => ./../../
