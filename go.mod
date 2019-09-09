module github.com/jinmukeji/plat-pkg

go 1.13

require (
	github.com/chzyer/logex v1.1.10 // indirect
	github.com/chzyer/test v0.0.0-20180213035817-a1ea475d72b1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible // indirect
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/hashicorp/consul/api v1.2.0
	github.com/jinmukeji/go-pkg v0.0.0-20190908165241-d2e61a4295a0
	github.com/jinmukeji/proto v0.0.0-20190909052444-48ae1aa3c845
	github.com/jinzhu/gorm v1.9.10
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.9.1
	github.com/micro/go-plugins v1.2.0
	github.com/micro/micro v1.9.1
	github.com/nats-io/nats-server/v2 v2.0.4 // indirect
	github.com/nlopes/slack v0.6.0 // indirect
	github.com/onsi/ginkgo v1.10.1 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	gopkg.in/yaml.v2 v2.2.2
)

// TODO: fix go mod tidy
replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.0
