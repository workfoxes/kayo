package indicator

// NewInterface creates a new Interface indicator.
type TradeIndicator interface {
	Initialize()
	Plot()
}
