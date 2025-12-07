package models

import "github.com/terrytay/nq/utils"

type Strategy interface {
	Execute([]utils.PriceData) BacktestResult
}

type BacktestResult struct {
	TotalProfit   float64
	WinRate       float64
	Trades        int
	WinningTrades int
}
