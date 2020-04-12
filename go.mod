module github.com/jinmukeji/plat-pkg/v2

go 1.14

replace (
	github.com/micro/go-micro/v2 => github.com/micro/go-micro/v2 v2.4.0
	// TODO: 修复 go-micro v2.4.0 相关包引用的 bug
	github.com/micro/micro/v2 => github.com/micro/micro/v2 v2.4.0

	// TODO: smallstep/cli v0.14.2 的 GOSUM 校验失败
	github.com/smallstep/cli => github.com/smallstep/cli v0.13.3
// go.etcd.io/etcd => go.etcd.io/etcd v0.0.0-20200401174654-e694b7bb0875
// github.com/coreos/etcd => go.etcd.io/etcd v0.0.0-20200401174654-e694b7bb0875
// github.com/coreos/etcd v3.3.18+incompatible => go.etcd.io/etcd v3.3.20+incompatible
)

require (
	github.com/coreos/etcd v3.3.20+incompatible // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gobwas/glob v0.2.3
	github.com/jinmukeji/go-pkg/v2 v2.2.7
	github.com/jinzhu/gorm v1.9.12
	github.com/manifoldco/promptui v0.7.0 // indirect
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.4.0
	github.com/micro/go-plugins/logger/logrus/v2 v2.3.0
	github.com/micro/go-plugins/micro/cors/v2 v2.3.0
	github.com/micro/go-plugins/micro/metadata/v2 v2.3.0
	github.com/micro/go-plugins/wrapper/service/v2 v2.3.0
	github.com/micro/micro/v2 v2.4.0
	github.com/prometheus/client_golang v1.5.1 // indirect
	github.com/samfoo/ansi v0.0.0-20160124022901-b6bd2ded7189 // indirect
	github.com/sirupsen/logrus v1.5.0
	github.com/smallstep/assert v0.0.0-20200103212524-b99dc1097b15 // indirect
	github.com/smallstep/cli v0.13.3
	github.com/stretchr/testify v1.5.1
	go.etcd.io/etcd v3.3.20+incompatible
	go.uber.org/zap v1.14.1 // indirect
	golang.org/x/crypto v0.0.0-20200406173513-056763e48d71 // indirect
	google.golang.org/grpc v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.2.8
)
