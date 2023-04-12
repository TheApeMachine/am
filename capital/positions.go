package capital

import (
	"encoding/json"
	"time"

	"github.com/wrk-grp/errnie"
	"github.com/wrk-grp/please"
)

type MarketPresenter struct {
	InstrumentName           string    `json:"instrumentName"`
	Expiry                   string    `json:"expiry"`
	MarketStatus             string    `json:"marketStatus"`
	Epic                     string    `json:"epic"`
	InstrumentType           string    `json:"instrumentType"`
	LotSize                  int       `json:"lotSize"`
	High                     float64   `json:"high"`
	Low                      float64   `json:"low"`
	PercentageChange         float64   `json:"percentageChange"`
	NetChange                float64   `json:"netChange"`
	Bid                      float64   `json:"bid"`
	Offer                    float64   `json:"offer"`
	UpdateTime               time.Time `json:"updateTime"`
	UpdateTimeUTC            time.Time `json:"updateTimeUTC"`
	DelayTime                int       `json:"delayTime"`
	StreamingPricesAvailable bool      `json:"streamingPricesAvailable"`
	ScalingFactor            int       `json:"scalingFactor"`
}

type PositionPresenter struct {
	Market   MarketPresenter `json:"market"`
	Position struct {
		ContractSize   int       `json:contractSize""`
		CreatedDate    time.Time `json:"createdDate"`
		CreatedDateUTC time.Time `json:"createdDateUTC"`
		DealID         string    `json:"dealId"`
		DealReference  string    `json:"dealReference"`
		WorkingOrderID string    `json:"workingOrderId"`
		Size           int       `json:"size"`
		Leverage       int       `json:"leverage"`
		UPL            float64   `json:"upl"`
		Direction      string    `json:"direction"`
		Level          float64   `json:"level"`
		Currency       string    `json:"currency"`
		GuaranteedStop bool      `json:"guaranteedStop"`
	} `json:"position"`
}

type Positions struct {
	Positions []PositionPresenter `positions`
}

type Position struct {
	client  *please.Request
	session SessionPresenter
}

func NewPosition(session SessionPresenter) *Position {
	return &Position{NewClient().conn, session}
}

func (position *Position) All() []PositionPresenter {
	position.client.AddHeaders(map[string]string{
		"CST":              position.session.CST,
		"X-SECURITY-TOKEN": position.session.Token,
	})

	res := position.client.Get("/api/v1/positions", map[string]interface{}{})

	data := Positions{}
	errnie.Handles(json.Unmarshal(res, &data))

	return data.Positions
}

func (position *Position) Enter(
	epic, direction string,
	size int,
	stoploss, takeprofit float64,
	guaranteedStop bool,
) error {
	position.client.AddHeaders(map[string]string{
		"CST":              position.session.CST,
		"X-SECURITY-TOKEN": position.session.Token,
	})

	res := position.client.Post("/api/v1/positions", map[string]interface{}{
		"epic":           epic,
		"direction":      direction,
		"size":           size,
		"stopAmount":     stoploss,
		"profitAmount":   takeprofit,
		"guaranteedStop": guaranteedStop,
	})

	data := Positions{}
	return errnie.Handles(json.Unmarshal(res, &data))
}

func (position *Position) Adjust(dealID string, stoploss, takeprofit float64) error {
	position.client.AddHeaders(map[string]string{
		"CST":              position.session.CST,
		"X-SECURITY-TOKEN": position.session.Token,
	})

	res := position.client.Put("/api/v1/positions", map[string]interface{}{
		"stopAmount":   stoploss,
		"profitAmount": takeprofit,
	})

	data := Positions{}
	return errnie.Handles(json.Unmarshal(res, &data))
}
