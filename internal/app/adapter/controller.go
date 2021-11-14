package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/orensimple/trade-order-app/internal/app/adapter/mysql"
	"github.com/orensimple/trade-order-app/internal/app/adapter/repository"
	"github.com/orensimple/trade-order-app/internal/app/adapter/service"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

// Controller is a controller
type Controller struct {
	OrderRepository repository.Order
	BillingService  service.Billing
}

// Router is routing settings
func Router() *gin.Engine {
	r := gin.Default()
	db := mysql.Connection()

	// init prometheus metrics
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	m.Use(r)

	orderRepository := repository.NewOrderRepo(db)

	ctrl := Controller{
		OrderRepository: orderRepository,
	}

	go ProcessOrder(ctrl)

	r.GET("/health", ctrl.health)

	api := r.Group("/api")
	api.POST("/order", ctrl.orderCreate)
	api.GET("/order/:id", ctrl.orderGet)
	api.PUT("/order/:id", ctrl.orderUpdate)
	api.DELETE("/order/:id", ctrl.orderDelete)

	return r
}
