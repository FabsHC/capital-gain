package services_test

import (
	"capital-gain/internal/config"
	"capital-gain/internal/models"
	"capital-gain/internal/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTaxCalculation(t *testing.T) {
	t.Parallel()
	reg := config.NewRegister()

	t.Run("should not pay taxes, no profit losses or gains", func(t *testing.T) {
		buyOperation := models.NewCapitalGainInput(models.BUY_OPERATION, 10, 100)
		sellOperation := models.NewCapitalGainInput(models.SELL_OPERATION, 10, 100)
		var input []models.CapitalGainInput
		input = append(input, *buyOperation)
		input = append(input, *sellOperation)

		taxCalculation := services.NewTaxCalculation(reg.BuyOperation, reg.SellOperation)

		outputList := taxCalculation.Execute(input)
		assert.Equal(t, 0.0, *outputList[0].Tax)
		assert.Equal(t, 0.0, *outputList[1].Tax)
	})

	t.Run("should not pay taxes, because operation has only profit losses", func(t *testing.T) {
		buyOperation := models.NewCapitalGainInput(models.BUY_OPERATION, 10, 100)
		sellOperation := models.NewCapitalGainInput(models.SELL_OPERATION, 3, 100)
		var input []models.CapitalGainInput
		input = append(input, *buyOperation)
		input = append(input, *sellOperation)

		taxCalculation := services.NewTaxCalculation(reg.BuyOperation, reg.SellOperation)

		outputList := taxCalculation.Execute(input)
		assert.Equal(t, 0.0, *outputList[0].Tax)
		assert.Equal(t, 0.0, *outputList[1].Tax)
	})

	t.Run("should not pay taxes, because sell operation value is lower than 20000", func(t *testing.T) {
		buyOperation := models.NewCapitalGainInput(models.BUY_OPERATION, 10, 100)
		sellOperation := models.NewCapitalGainInput(models.SELL_OPERATION, 15, 30)
		var input []models.CapitalGainInput
		input = append(input, *buyOperation)
		input = append(input, *sellOperation)

		taxCalculation := services.NewTaxCalculation(reg.BuyOperation, reg.SellOperation)

		outputList := taxCalculation.Execute(input)
		assert.Equal(t, 0.0, *outputList[0].Tax)
		assert.Equal(t, 0.0, *outputList[1].Tax)
	})

	t.Run("should pay taxes, because sell operation value is lower than 20000", func(t *testing.T) {
		buyOperation := models.NewCapitalGainInput(models.BUY_OPERATION, 10, 100)
		sellOperation := models.NewCapitalGainInput(models.SELL_OPERATION, 350, 100)
		var input []models.CapitalGainInput
		input = append(input, *buyOperation)
		input = append(input, *sellOperation)

		taxCalculation := services.NewTaxCalculation(reg.BuyOperation, reg.SellOperation)

		outputList := taxCalculation.Execute(input)
		assert.Equal(t, 0.0, *outputList[0].Tax)
		assert.Equal(t, 6800.0, *outputList[1].Tax)
	})

	t.Run("should not pay taxes, because profit gains will cover the losses", func(t *testing.T) {
		buyOperation := models.NewCapitalGainInput(models.BUY_OPERATION, 10, 10000)
		sellOperation := models.NewCapitalGainInput(models.SELL_OPERATION, 2, 5000)
		var input []models.CapitalGainInput
		input = append(input, *buyOperation)
		input = append(input, *sellOperation)
		sellOperation = models.NewCapitalGainInput(models.SELL_OPERATION, 20, 2000)
		input = append(input, *sellOperation)
		sellOperation = models.NewCapitalGainInput(models.SELL_OPERATION, 20, 2000)
		input = append(input, *sellOperation)

		taxCalculation := services.NewTaxCalculation(reg.BuyOperation, reg.SellOperation)

		outputList := taxCalculation.Execute(input)
		assert.Equal(t, 0.0, *outputList[0].Tax)
		assert.Equal(t, 0.0, *outputList[1].Tax)
		assert.Equal(t, 0.0, *outputList[2].Tax)
		assert.Equal(t, 0.0, *outputList[3].Tax)
	})
}
