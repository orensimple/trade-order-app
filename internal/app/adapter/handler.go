package adapter

import (
	"net/http"

	"github.com/prometheus/common/log"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/orensimple/trade-order-app/internal/app/application/usecase"
	"github.com/orensimple/trade-order-app/internal/app/domain"
)

func (ctrl Controller) health(c *gin.Context) {
	c.JSON(http.StatusOK, domain.SimpleResponse{Status: "OK"})
}

func (ctrl Controller) orderCreate(c *gin.Context) {
	req := new(domain.CreateOrderRequest)
	err := render.Bind(c.Request, req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong request params"})

		return
	}

	order := &domain.Order{
		ID:           uuid.New(),
		AccountID:    req.AccountID,
		InstrumentID: req.InstrumentID,
		Type:         req.Type,
		Price:        req.Price,
		Volume:       req.Volume,
		Status:       "new",
	}

	err = usecase.BillingBlocked(ctrl.BillingService, order)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusConflict, domain.SimpleResponse{Status: "fail blocked money"})

		return
	}

	order, err = usecase.CreateOrder(ctrl.OrderRepository, order)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, domain.SimpleResponse{Status: "failed create order"})

		return
	}

	c.JSON(http.StatusOK, order)
}

func (ctrl Controller) orderGet(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong id"})

		return
	}

	res, err := usecase.GetOrder(ctrl.OrderRepository, &domain.Order{ID: id})
	if err != nil && err.Error() != "order not found" {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "failed get order"})

		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, domain.SimpleResponse{Status: "order not found"})

		return
	}

	c.JSON(http.StatusOK, res)
}

func (ctrl Controller) orderUpdate(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong id"})

		return
	}

	order, err := usecase.GetOrder(ctrl.OrderRepository, &domain.Order{ID: id})
	if err != nil && err.Error() != "order not found" {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "failed get order"})

		return
	}
	if order == nil {
		c.JSON(http.StatusNotFound, domain.SimpleResponse{Status: "order not found"})

		return
	}

	var req domain.CreateOrderRequest
	err = c.ShouldBind(&req)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong request params"})

		return
	}

	err = usecase.UpdateOrder(ctrl.OrderRepository, order)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, domain.SimpleResponse{Status: "failed delete order"})

		return
	}

	c.JSON(http.StatusOK, order)
}

func (ctrl Controller) orderDelete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "wrong id"})

		return
	}

	order, err := usecase.GetOrder(ctrl.OrderRepository, &domain.Order{ID: id})
	if err != nil && err.Error() != "order not found" {
		log.Error(err)
		c.JSON(http.StatusBadRequest, domain.SimpleResponse{Status: "failed get order"})

		return
	}
	if order == nil {
		c.JSON(http.StatusNotFound, domain.SimpleResponse{Status: "order not found"})

		return
	}

	order.Volume = 0 - order.Volume
	err = usecase.BillingBlocked(ctrl.BillingService, order)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusConflict, domain.SimpleResponse{Status: "fail unblocked money"})

		return
	}

	order.Status = "canceled"
	order.Volume = 0 - order.Volume
	err = usecase.UpdateOrder(ctrl.OrderRepository, order)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, domain.SimpleResponse{Status: "failed delete order"})

		return
	}

	c.JSON(http.StatusOK, domain.SimpleResponse{Status: "OK"})
}
