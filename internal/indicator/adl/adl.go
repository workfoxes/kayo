package adl

import (
	commonB "github.com/workfoxes/kayo/internal/broker/common"
	"github.com/workfoxes/kayo/internal/indicator/common"
	"github.com/workfoxes/kayo/internal/model"
)

type Indicator struct {
	common.BaseIndicator
}

func (i *Indicator) Process(item *commonB.Item) bool {
	return false
}

func (i *Indicator) SetFilterCheck(check model.FilterCheck) error {
	return nil
}
