module github.com/ZQCard/kratos-base-project

go 1.16

require (
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20220414054820-d0b704b8f38d
	github.com/go-kratos/kratos/v2 v2.2.1
	github.com/go-redis/redis/v8 v8.11.5
	github.com/golang-jwt/jwt/v4 v4.4.1
	github.com/google/wire v0.5.0
	github.com/gorilla/handlers v1.5.1
	github.com/hashicorp/consul/api v1.12.0
	go.opentelemetry.io/otel v1.6.3
	go.opentelemetry.io/otel/exporters/jaeger v1.6.3
	go.opentelemetry.io/otel/sdk v1.6.3
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/genproto v0.0.0-20220421151946-72621c1f0bd3
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.4
)

require (
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.3.7 // indirect
)
