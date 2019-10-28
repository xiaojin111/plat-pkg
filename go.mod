module github.com/jinmukeji/plat-pkg

go 1.13

// TODO: fix go mod tidy
replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.0

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gobwas/glob v0.2.3
	github.com/jinmukeji/go-pkg v0.0.0-20191028134007-b15a5e55a35b
	github.com/jinmukeji/proto v0.0.0-20191027101206-85d22e6cfe99
	github.com/jinzhu/gorm v1.9.11
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.14.0
	github.com/micro/go-plugins v1.4.0
	github.com/micro/micro v1.14.0
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	go.etcd.io/etcd v3.3.17+incompatible
	gopkg.in/yaml.v2 v2.2.4
)
