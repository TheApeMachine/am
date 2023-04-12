package capital

import (
	"encoding/json"
	"net/url"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/websocket"
	"github.com/theapemachine/wrkspc/tweaker"
	"github.com/wrk-grp/errnie"
)

type MarketData struct {
	Status      string `json:"status"`
	Destination string `json:"destination"`
	Payload     struct {
		Epic      string  `json:"epic"`
		Product   string  `json:"product"`
		Bid       float64 `json:"bid"`
		BidQty    float64 `json:"bidQty"`
		Offer     float64 `json:"ofr"`
		OfferQty  float64 `json:"offerQty"`
		Timestamp int64   `json:"timestamp"`
	} `json:"payload"`
}

type OHLCData struct {
	Status      string `json:"status"`
	Destination string `json:"destination"`
	Payload     struct {
		Resolution string  `json:"resolution"`
		Epic       string  `json:"epic"`
		Type       string  `json:"type"`
		PriceType  string  `json:"priceType"`
		Timestamp  int64   `json:"t"`
		High       float64 `json:"h"`
		Low        float64 `json:"l"`
		Open       float64 `json:"open"`
		Close      float64 `json:"close"`
	} `json:"payload"`
}

type Stream struct {
	conn    *websocket.Conn
	session SessionPresenter
}

func NewStream(session SessionPresenter) *Stream {
	u := url.URL{
		Scheme: "wss", Host: tweaker.GetString("capital.stream"), Path: "/connect",
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	errnie.Handles(err)

	return &Stream{
		conn:    conn,
		session: session,
	}
}

func (stream *Stream) Read() chan OHLCData {
	out := make(chan OHLCData)

	go func() {
		defer close(out)

		for {
			_, msg, err := stream.conn.ReadMessage()
			errnie.Handles(err)

			spew.Dump(msg)

			marketData := OHLCData{}
			errnie.Handles(json.Unmarshal(msg, &marketData))

			out <- marketData
		}
	}()

	return out
}

func (stream *Stream) Write(epics ...string) {
	data := map[string]interface{}{
		"destination":   "OHLCMarketData.subscribe",
		"correlationId": "3",
		"cst":           stream.session.CST,
		"securityToken": stream.session.Token,
		"payload": map[string]interface{}{
			"epics":       epics,
			"resolutions": []string{"MINUTE_1"},
			"type":        "classic",
		},
	}

	buf, err := json.Marshal(data)
	errnie.Handles(err)

	errnie.Handles(
		stream.conn.WriteMessage(websocket.TextMessage, buf),
	)
}
