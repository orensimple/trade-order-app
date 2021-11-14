module github.com/orensimple/trade-order-app

// +heroku goVersion go1.15
go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-chi/render v1.0.1 // indirect
	github.com/google/uuid v1.1.2
	github.com/penglongli/gin-metrics v0.1.4
	github.com/prometheus/common v0.10.0
	github.com/shopspring/decimal v1.3.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.5.1 // indirect
	gorm.io/driver/mysql v1.1.1
	gorm.io/gorm v1.21.9
	gorm.io/plugin/dbresolver v1.1.0
)
