package indicator

import (
	"github.com/workfoxes/kayo/internal/broker/common"
	"github.com/workfoxes/kayo/internal/model"
)

// TradeIndicator NewInterface creates a new Interface indicator.
type TradeIndicator interface {
	ID() string
	Initialize()
	Plot()
	Process(item *common.Item) bool
	SetFilterCheck(check model.FilterCheck) error
}
