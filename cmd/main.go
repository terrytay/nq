package main

import (
	"fmt"
	"sync"

	"github.com/terrytay/nq/models"
	"github.com/terrytay/nq/strategies"
	"github.com/terrytay/nq/utils"
)

func main() {
	data, err := utils.LoadHistoricalData("data.csv")
	if err != nil {
		panic(err)
	}
	strategies := []models.Strategy{
		strategies.MovingAverageStrategy{ShortPeriod: 10, LongPeriod: 20},
		strategies.MovingAverageStrategy{ShortPeriod: 20, LongPeriod: 50},
	}

	var wg sync.WaitGroup
	results := make(chan models.BacktestResult)

	for _, strategy := range strategies {
		wg.Go(func() {
			results <- strategy.Execute(data)
		})
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("Profit: %.2f, Win Rate: %.2f, Trades %d\n", result.TotalProfit, result.WinRate, result.Trades)
	}

}
