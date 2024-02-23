package util

import "math"

func CalculateAveragePrice(totalStocks, stocksPurchased uint, actualAveragePrice, purchasePrice float64) float64 {
	totalStocksFloat := float64(totalStocks)
	stocksPurchasedFloat := float64(stocksPurchased)
	average := ((totalStocksFloat * actualAveragePrice) + (stocksPurchasedFloat * purchasePrice)) / (totalStocksFloat + stocksPurchasedFloat)
	return RoundTwoDecimals(average, 2)
}

func RoundTwoDecimals(value float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(value*ratio) / ratio
}
