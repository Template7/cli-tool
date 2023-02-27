module cli-tool

go 1.16

require (
	github.com/Pallinder/go-randomdata v1.2.0
	github.com/Template7/backend v1.2.2-0.20220126104422-13caa54e118b
	github.com/Template7/common v0.0.0-20230227182610-f3dec298b5dc
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/klauspost/compress v1.14.2 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.10.1
	github.com/urfave/cli/v2 v2.3.0
	go.mongodb.org/mongo-driver v1.8.2
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce
	golang.org/x/net v0.0.0-20220121210141-e204ce36a2ba // indirect
	gopkg.in/ini.v1 v1.66.3 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	gorm.io/driver/mysql v1.2.3
	gorm.io/gorm v1.22.5
)

//replace (
//	github.com/Template7/common => ../common
//)
