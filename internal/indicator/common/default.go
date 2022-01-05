package common

import (
	"github.com/workfoxes/kayo/internal/broker/common"
	"github.com/workfoxes/kayo/internal/model"
)

type Signal int

// BaseIndicator : Base Indicator for all the indicator support in kayo
type BaseIndicator struct {
	Name      string
	ShortName string
	Signal    Signal
}

// SupportedIndicators  - list all supporting trading indicator
const (
	RSI                          = "RSI"
	MACD                         = "MACD"
	StochasticOscillator         = "StochasticOscillator"
	AverageDirectionalIndex      = "AverageDirectionalIndex"
	OnBalanceVolume              = "OnBalanceVolume"
	AccumulationDistributionLine = "AccumulationDistributionLine"
)

func (i *BaseIndicator) Initialize() {
	i.Name = "Name"
	i.ShortName = "ShortName"
}

func (i *BaseIndicator) Plot() {
	i.Name = "Name"
	i.ShortName = "ShortName"
}

func (i *BaseIndicator) ID() string {
	return i.Name
}

func (i *BaseIndicator) Process(item *common.Item) {
	i.Name = "Name"
	i.ShortName = "ShortName"
}

func (i *BaseIndicator) SetFilterCheck(check model.FilterCheck) {
	i.Name = check.Indicator
}
