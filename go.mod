module github.com/jinmukeji/plat-pkg

go 1.13

// TODO: fix go mod tidy
replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.0

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gobwas/glob v0.2.3
	github.com/hashicorp/consul/api v1.1.0
	github.com/jinmukeji/go-pkg v0.0.0-20190914100229-9bbbe1b3a1b2
	github.com/jinmukeji/proto v0.0.0-20190914101010-394a4c90ecf1
	github.com/jinzhu/gorm v1.9.10
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.10.0
	github.com/micro/go-plugins v1.3.0
	github.com/micro/micro v1.10.0
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	gopkg.in/yaml.v2 v2.2.2
)
