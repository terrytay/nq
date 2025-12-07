package strategies

import (
	"github.com/terrytay/nq/models"
	"github.com/terrytay/nq/utils"
)

type MovingAverageStrategy struct {
	ShortPeriod int
	LongPeriod  int
}

func (s MovingAverageStrategy) Execute(data []utils.PriceData) models.BacktestResult {
	var result models.BacktestResult
	var inPosition bool
	var entryPrice float64

	for i := s.LongPeriod; i < len(data); i++ {
		shortMa := calculateMA(data[i-s.ShortPeriod:i], s.ShortPeriod)
		longMa := calculateMA(data[i-s.LongPeriod:i], s.LongPeriod)

		if shortMa > longMa && !inPosition {
			entryPrice = data[i].Close
			inPosition = true
		} else if shortMa < longMa && inPosition {
			exitPrice := data[i].Close
			profit := exitPrice - entryPrice
			result.TotalProfit += profit
			result.Trades++
			if profit > 0 {
				result.WinningTrades++
			}
			inPosition = false
			// fmt.Printf("Entry: %.2f, Exit: %.2f, Profit: %.2f\n", entryPrice, exitPrice, profit)
		}
	}

	if result.Trades > 0 {
		result.WinRate = float64(result.WinningTrades) / float64(result.Trades)
	}
	return result
}

func calculateMA(data []utils.PriceData, period int) float64 {
	var sum float64
	for _, price := range data {
		sum += price.Close
	}

	return sum / float64(period)
}
