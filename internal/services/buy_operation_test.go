package services_test

import (
	"capital-gain/internal/models"
	"capital-gain/internal/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuyOperation(t *testing.T) {
	t.Parallel()
	buyOperationService := services.NewBuyOperation()
	purchase := models.NewStocksInfo()

	t.Run("should execute buy operation and set average price to 10 on first execution", func(t *testing.T) {
		operation := models.NewCapitalGainInput(models.BUY_OPERATION, 10, 100)
		buyOperationService.Execute(purchase, *operation)

		assert.Equal(t, float64(10), purchase.AveragePrice)
		assert.Equal(t, operation.Quantity, purchase.Shares)
	})

	t.Run("should execute buy operation and set average price to 7.5 and 200 total stocks on second execution", func(t *testing.T) {
		operation := models.NewCapitalGainInput(models.BUY_OPERATION, 5, 100)
		buyOperationService.Execute(purchase, *operation)

		assert.Equal(t, 7.5, purchase.AveragePrice)
		assert.Equal(t, uint(200), purchase.Shares)
	})
}
