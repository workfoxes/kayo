package indicator

import (
	"reflect"
	"testing"

	"github.com/workfoxes/kayo/internal/indicator/common"
	"github.com/workfoxes/kayo/internal/indicator/macd"
	"github.com/workfoxes/kayo/internal/indicator/rsi"
	"github.com/workfoxes/kayo/internal/indicator/so"
)

func TestIndicatorPositive(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want TradeIndicator
	}{
		{name: "RSI_positive", args: args{name: common.RSI}, want: new(rsi.RSI)},
		{name: "MACD_positive", args: args{name: common.MACD}, want: new(macd.Indicator)},
		{name: "SO_positive", args: args{name: common.StochasticOscillator}, want: new(so.Indicator)},
		// {name: "IBPosition", args: args{brokerName: _default.InteractiveBroker}, want: new(ib.InteractiveBroker)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIndicator(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIndicator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndicatorNegative(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "RSI_positive", args: args{name: "RSI1"}},
		{name: "MACD_positive", args: args{name: "MACD1"}},
		{name: "SO_positive", args: args{name: "StochasticOscillator1"}},
		// {name: "IBPosition", args: args{brokerName: _default.InteractiveBroker}, want: new(ib.InteractiveBroker)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIndicator(tt.args.name); got != nil {
				t.Errorf("NewIndicator() = %v, want nil", got)
			}
		})
	}
}
