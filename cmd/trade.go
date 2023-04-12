package cmd

import (
	"github.com/spf13/cobra"
	"github.com/theapemachine/am/capital"
	"github.com/wrk-grp/errnie"
)

/*
tradeCmd starts a capital.com trading bot.
*/
var tradeCmd = &cobra.Command{
	Use:   "trade",
	Short: "Start the trading bot.",
	Long:  tradetxt,
	RunE: func(_ *cobra.Command, _ []string) (err error) {
		errnie.Trace()

		session := capital.NewSession().Open()
		// positions := capital.NewPosition(session).All()
		// markets := capital.NewMarket(session).Details("SILVER")

		stream := capital.NewStream(session)
		trades := map[string]*capital.Trade{
			"SILVER": capital.NewTrade("SILVER", session),
		}

		go func() {
			for signal := range stream.Read() {
				trades[signal.Payload.Epic].Tick(signal)
			}
		}()

		stream.Write("SILVER")

		select {}
		return nil
	},
}

/*
tradetxt lives here to keep the command definition section cleaner.
*/
var tradetxt = `
Run the trading bot.
`
