package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/shopspring/decimal"

	"github.com/orensimple/trade-order-app/internal/app/domain"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
)

type billingPayRequest struct {
	PayAmount decimal.Decimal `json:"pay_amount"`
}
type billingBlockedRequest struct {
	BlockedAmount decimal.Decimal `json:"blocked_amount"`
}

// Billing is billing app
type Billing struct{}

func (b Billing) Pay(o *domain.Order) error {
	host := viper.Get("billing_host")
	url := fmt.Sprintf("%v/api/account/%s/pay", host, o.AccountID)

	data, err := json.Marshal(billingPayRequest{PayAmount: o.Price.Mul(decimal.NewFromInt(o.Volume))})
	if err != nil {
		return err
	}

	log.Info("Sending request")
	res, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	log.Infof("Dump response with code, '%d'", res.StatusCode)
	dump, err := httputil.DumpResponse(res, true)
	if err == nil {
		log.Debugf("billing response '%q", dump)
	}
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New("something wrong")
	}

	return nil
}

func (b Billing) Blocked(o *domain.Order) error {
	host := viper.Get("billing_host")
	url := fmt.Sprintf("%v/api/account/%s/block", host, o.AccountID)

	data, err := json.Marshal(billingBlockedRequest{BlockedAmount: o.Price.Mul(decimal.NewFromInt(o.Volume))})
	if err != nil {
		return err
	}

	log.Info("Sending request")
	res, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	log.Infof("Dump response with code, '%d'", res.StatusCode)
	dump, err := httputil.DumpResponse(res, true)
	if err == nil {
		log.Debugf("billing response '%q", dump)
	}
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New("something wrong")
	}

	return nil
}
