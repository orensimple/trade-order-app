package adapter

import (
	"math/rand"
	"time"

	"github.com/orensimple/trade-order-app/internal/app/application/usecase"
	"github.com/orensimple/trade-order-app/internal/app/domain"
	"github.com/prometheus/common/log"
)

func ProcessOrder(ctrl Controller) {
	for {
		order, err := usecase.GetOrder(ctrl.OrderRepository, &domain.Order{Status: "new"})
		if err != nil && err.Error() != "order not found" {
			log.Error(err)
		}

		if order == nil {
			time.Sleep(10 * time.Second)

			continue
		}

		if rand.Intn(2) == 0 {
			err := usecase.BillingPay(ctrl.BillingService, order)
			if err != nil {
				log.Error(err)
				continue
			}

			order.Status = "done"
			err = usecase.UpdateOrder(ctrl.OrderRepository, order)
			if err != nil {
				log.Error(err)
			}

			log.Info("done")
			continue
		}

		err = usecase.BillingBlocked(ctrl.BillingService, order)
		if err != nil {
			log.Error(err)

			continue
		}

		order.Status = "fail"
		order.Volume = 0 - order.Volume
		err = usecase.UpdateOrder(ctrl.OrderRepository, order)
		if err != nil {
			log.Error(err)
		}

		log.Info("fail")
	}
}
