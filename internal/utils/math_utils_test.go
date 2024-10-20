package utils_test

import (
	"capital-gain/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMathUtils(t *testing.T) {
	t.Parallel()

	t.Run("should execute CalculateAverageSharePrice with success", func(t *testing.T) {
		var (
			totalStocks        = uint(0)
			stocksPurchased    = uint(100)
			actualAveragePrice = 0.0
			purchasePrice      = 10.0
		)
		value := utils.CalculateAverageSharePrice(totalStocks, stocksPurchased, actualAveragePrice, purchasePrice)
		assert.Equal(t, 10.0, value)

		totalStocks = uint(100)
		actualAveragePrice = 10.0
		purchasePrice = 5.0
		value = utils.CalculateAverageSharePrice(totalStocks, stocksPurchased, actualAveragePrice, purchasePrice)
		assert.Equal(t, 7.5, value)
	})

	t.Run("should execute RoundTwoDecimals with success", func(t *testing.T) {
		value := utils.RoundTwoDecimals(15.379, 2)
		assert.Equal(t, 15.38, value)
	})
}
