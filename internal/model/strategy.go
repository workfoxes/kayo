package model

type Strategy struct {
	SystemModel
	Name            string        `gorm:"column:Name;" json:"Name"`
	Description     string        `gorm:"column:Description;default:null" json:"Description"`
	TakeProfit      ItemPointer   `json:"TakeProfit,omitempty" gorm:"foreignKey:StrategyID;references:id"`
	StopLoss        ItemPointer   `json:"StopLoss,omitempty" gorm:"foreignKey:StrategyID;references:id"`
	BuyFilterCheck  []FilterCheck `json:"BuyFilterCheck,omitempty" gorm:"foreignKey:StrategyID;references:id"`
	SellFilterCheck []FilterCheck `json:"SellFilterCheck,omitempty" gorm:"foreignKey:StrategyID;references:id"`

	//BuyFilters  []*indicator.TradeIndicator `gorm:"-" json:"BuyFilters"`
	//SellFilters []*indicator.TradeIndicator `gorm:"-" json:"SellFilters"`
}

type ItemPointer struct {
	SystemModel
	Type       string  `json:"Type"`
	Value      float64 `json:"Value"`
	StrategyID int     `json:"StrategyID"`
}
type IndicatorParams struct {
	SystemModel
	Key       string `gorm:"column:Key;" json:"Key"`
	Value     string `gorm:"column:Value;" json:"Value"`
	CheckerID int    `json:"CheckerID"`
}
type FilterCheck struct {
	SystemModel
	Indicator       string            `json:"Indicator"`
	IndicatorParams []IndicatorParams `json:"IndicatorParams,omitempty" gorm:"foreignKey:CheckerID;references:id"`
	StrategyID      int               `gorm:"column:StrategyID;" json:"StrategyID"`
}

func (FilterCheck) TableName() string {
	return "filter_check"
}

// TableName overrides the table name used by User to `profiles`
func (IndicatorParams) TableName() string {
	return "indicator_param"
}

// TableName overrides the table name used by User to `profiles`
func (ItemPointer) TableName() string {
	return "item_pointer"
}

// TableName overrides the table name used by User to `profiles`
func (Strategy) TableName() string {
	return "strategy"
}
