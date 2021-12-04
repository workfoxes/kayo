package indicator

import "github.com/workfoxes/kayo/internal/broker/common"

// TradeIndicator NewInterface creates a new Interface indicator.
type TradeIndicator interface {
	Initialize()
	Plot()
	Process(item *common.Item)
}
