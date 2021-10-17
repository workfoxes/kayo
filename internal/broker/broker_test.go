package broker

import (
	"github.com/workfoxes/kayo/internal/broker/binance"
	"github.com/workfoxes/kayo/internal/broker/common"
	"github.com/workfoxes/kayo/internal/broker/ib"
	"reflect"
	"testing"
)

func TestNewBroker(t *testing.T) {
	type args struct {
		brokerName string
	}
	tests := []struct {
		name string
		args args
		want StockBroker
	}{
		{name: "BinancePosition", args: args{brokerName: common.Binance}, want: new(binance.Binance)},
		{name: "IBPosition", args: args{brokerName: common.InteractiveBroker}, want: new(ib.InteractiveBroker)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBroker(tt.args.brokerName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBroker() = %v, want %v", got, tt.want)
			}
		})
	}
}
