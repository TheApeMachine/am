package capital

import "github.com/davecgh/go-spew/spew"

type Trade struct {
	Symbol     string
	Confidence float64
	Price      float64
	Profit     float64
	Active     *bool
	session    SessionPresenter
}

func NewTrade(symbol string, session SessionPresenter) *Trade {
	return &Trade{
		Symbol:  symbol,
		session: session,
	}
}

func (trade *Trade) Tick(data OHLCData) {
	switch data.Destination {
	case "ohlc.event":
		spew.Dump(data)
	}
}

func (trade *Trade) Enter() {
	position := NewPosition(trade.session)
	position.Enter(trade.Symbol, "BUY", 1, trade.Price-0.01, trade.Price+0.01, true)
}

func (trade *Trade) Exit() {
	position := NewPosition(trade.session)
	position.Enter(trade.Symbol, "SELL", 1, trade.Price-0.01, trade.Price+0.01, true)
}
