package capital

import (
	"encoding/json"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/wrk-grp/errnie"
	"github.com/wrk-grp/please"
)

type Node struct {
	ID   string `id`
	Name string `name`
}

type TopLevel struct {
	Nodes []Node `json:"nodes"`
}

type MarketDetail struct {
	DelayTime                int       `json:"delayTime"`
	Epic                     string    `json:"epic"`
	NetChange                float64   `json:"netChange"`
	LotSize                  int       `json:"lotSize"`
	Expiry                   string    `json:"expiry"`
	InstrumentType           string    `json:"instrumentType"`
	InstrumentName           string    `json:"instrumentName"`
	High                     float64   `json:"high"`
	Low                      float64   `json:"low"`
	PercentageChange         float64   `json:"percentageChange"`
	UpdateTime               time.Time `json:"updateTime"`
	UpdateTimeUTC            time.Time `json:"updateTimeUTC"`
	Bid                      float64   `json:"bid"`
	Offer                    float64   `json:"offer"`
	StreamingPricesAvailable bool      `json:"streamingPricesAvailable"`
	MarketStatus             string    `json:"marketStatus"`
	ScalingFactor            int       `json:"scalingFactor"`
}

type MarketsPresenter struct {
	Markets []MarketDetail `json:"markets"`
}

type OpeningHours struct {
	Mon  []string `json:"mon"`
	Tue  []string `json:"tue"`
	Wed  []string `json:"wed"`
	Thu  []string `json:"thu"`
	Fri  []string `json:"fri"`
	Sat  []string `json:"sat"`
	Sun  []string `json:"sun"`
	Zone string   `json:"zone"`
}

type InstrumentPresenter struct {
	Epic                     string       `json:"epic"`
	Expiry                   string       `json:"expiry"`
	Name                     string       `json:"name"`
	LotSize                  int          `json:"lotSize"`
	Type                     string       `json:"type"`
	ControlledRiskAllowed    bool         `json:"controlledRiskAllowed"`
	StreamingPricesAvailable bool         `json:"streamingPricesAvailable"`
	Currency                 string       `json:"currency"`
	MarginFactor             int          `json:"marginFactor"`
	MarginFactorUnit         string       `json:"marginFactorUnit"`
	OpeningHours             OpeningHours `json:"openingHours"`
	Country                  string       `json:"country"`
}

type DealingRulesPresenter struct {
	MinStepDistance struct {
		Unit  string  `json:"unit"`
		Value float64 `json:"value"`
	} `json:"minStepDistance"`
	MinDealSize struct {
		Unit  string  `json:"unit"`
		Value float64 `json:"value"`
	} `json:"minDealSize"`
	MinControlledRiskStopDistance struct {
		Unit  string  `json:"unit"`
		Value float64 `json:"value"`
	} `json:"minControlledRiskStopDistance"`
	MinNormalStopOrLimitDistance struct {
		Unit  string  `json:"unit"`
		Value float64 `json:"value"`
	} `json:"minNormalStopOrLimitDistance"`
	MaxStopOrLimitDistance struct {
		Unit  string  `json:"unit"`
		Value float64 `json:"value"`
	} `json:"maxStopOrLimitDistance"`
	MarketOrderPreference   string `json:"marketOrderPreference"`
	TrailingStopsPreference string `json:"trailingStopsPreference"`
}

type SnapshotPresenter struct {
	MarketStatus        string    `json:"marketStatus"`
	NetChange           float64   `json:"netChange"`
	PercentageChange    float64   `json:"percentageChange"`
	UpdateTime          time.Time `json:"updateTime"`
	DelayTime           int       `json:"delayTime"`
	Bid                 float64   `json:"bid"`
	Offer               float64   `json:"offer"`
	High                float64   `json:"high"`
	Low                 float64   `json:"low"`
	DecimalPlacesFactor int       `json:"decimalPlacesFactor"`
	ScalingFactor       int       `json:"scalingFactor"`
}

type DetailPresenter struct {
	Instrument   InstrumentPresenter   `json:"instrument"`
	DealingRules DealingRulesPresenter `json:"dealingRules"`
	Snapshot     SnapshotPresenter     `json:"snapshot"`
}

type Market struct {
	client  *please.Request
	session SessionPresenter
}

func NewMarket(session SessionPresenter) *Market {
	return &Market{NewClient().conn, session}
}

func (market *Market) All() MarketsPresenter {
	market.client.AddHeaders(map[string]string{
		"CST":              market.session.CST,
		"X-SECURITY-TOKEN": market.session.Token,
	})

	res := market.client.Get("/api/v1/markets", map[string]interface{}{
		"epics": "SILVER",
	})

	spew.Dump(res)

	data := MarketsPresenter{}
	errnie.Handles(json.Unmarshal(res, &data))

	return data
}

func (market *Market) Details(epic string) DetailPresenter {
	market.client.AddHeaders(map[string]string{
		"CST":              market.session.CST,
		"X-SECURITY-TOKEN": market.session.Token,
	})

	res := market.client.Get("/api/v1/markets/"+epic, map[string]interface{}{})

	data := DetailPresenter{}
	errnie.Handles(json.Unmarshal(res, &data))

	return data
}
