package services_test

import (
	"capital-gain/internal/models"
	"capital-gain/internal/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSellOperation(t *testing.T) {
	t.Parallel()
	buyOperation := services.NewBuyOperation()
	sellOperation := services.NewSellOperation()

	t.Run("should execute sell operation and do not have any gains/losses because average price is same as selling price", func(t *testing.T) {
		purchase := models.NewStocksInfo()
		sale := models.NewProfit()
		operation := models.NewCapitalGainInput(models.BUY_OPERATION, 10, 100)

		buyOperation.Execute(purchase, *operation)
		assert.Equal(t, float64(10), purchase.AveragePrice)

		operation = models.NewCapitalGainInput(models.SELL_OPERATION, 10, 100)
		sellOperation.Execute(purchase, sale, *operation)
		assert.Equal(t, float64(0), sale.Gains)
		assert.Equal(t, float64(0), sale.Losses)
	})

	t.Run("should execute sell operation and have losses because selling price is bellow average price", func(t *testing.T) {
		purchase := models.NewStocksInfo()
		sale := models.NewProfit()
		operation := models.NewCapitalGainInput(models.BUY_OPERATION, 10, 100)

		buyOperation.Execute(purchase, *operation)
		assert.Equal(t, float64(10), purchase.AveragePrice)

		operation = models.NewCapitalGainInput(models.SELL_OPERATION, 5, 100)
		sellOperation.Execute(purchase, sale, *operation)
		assert.Equal(t, float64(0), sale.Gains)
		assert.Equal(t, float64(500), sale.Losses)
	})

	t.Run("should execute sell operation and have gains because selling price is higher than average price", func(t *testing.T) {
		purchase := models.NewStocksInfo()
		sale := models.NewProfit()
		operation := models.NewCapitalGainInput(models.BUY_OPERATION, 10, 100)

		buyOperation.Execute(purchase, *operation)
		assert.Equal(t, float64(10), purchase.AveragePrice)

		operation = models.NewCapitalGainInput(models.SELL_OPERATION, 25, 100)
		sellOperation.Execute(purchase, sale, *operation)
		assert.Equal(t, float64(1500), sale.Gains)
		assert.Equal(t, float64(0), sale.Losses)
	})
}
