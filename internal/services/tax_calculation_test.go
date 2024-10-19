package services_test

import (
	"capital-gain/internal/models"
	"capital-gain/internal/services"
	"testing"
)

func TestShouldNotPayTaxesNoProfitLossesOrGains(t *testing.T) {
	buyOperation := models.NewCapitalGainInput(models.BUY_OPERATION, 10, 100)
	sellOperation := models.NewCapitalGainInput(models.SELL_OPERATION, 10, 100)
	var input []models.CapitalGainInput
	input = append(input, *buyOperation)
	input = append(input, *sellOperation)

	taxCalculation := services.NewTaxCalculation(services.NewBuyOperation(), services.NewSellOperation())

	outputList := taxCalculation.Execute(input)
	if *outputList[0].Tax > 0 {
		t.Error("Tax calculation error, purchase operations do not pay taxes")
	}
	if *outputList[1].Tax > 0 {
		t.Error("Tax calculation error, sales operations without profits must not pay taxes")
	}
}

func TestShouldNotPayTaxesBecauseProfitLosses(t *testing.T) {
	buyOperation := models.NewCapitalGainInput(models.BUY_OPERATION, 10, 100)
	sellOperation := models.NewCapitalGainInput(models.SELL_OPERATION, 3, 100)
	var input []models.CapitalGainInput
	input = append(input, *buyOperation)
	input = append(input, *sellOperation)

	taxCalculation := services.NewTaxCalculation(services.NewBuyOperation(), services.NewSellOperation())

	outputList := taxCalculation.Execute(input)
	if *outputList[0].Tax > 0 {
		t.Error("Tax calculation error, purchase operations do not pay taxes")
	}
	if *outputList[1].Tax > 0 {
		t.Error("Tax calculation error, loss-making sales transactions should not pay taxes")
	}
}

func TestShouldNotPayTaxesBecauseSellOperationValueLowerThan20000(t *testing.T) {
	buyOperation := models.NewCapitalGainInput(models.BUY_OPERATION, 10, 100)
	sellOperation := models.NewCapitalGainInput(models.SELL_OPERATION, 15, 30)
	var input []models.CapitalGainInput
	input = append(input, *buyOperation)
	input = append(input, *sellOperation)

	taxCalculation := services.NewTaxCalculation(services.NewBuyOperation(), services.NewSellOperation())

	outputList := taxCalculation.Execute(input)
	if *outputList[0].Tax > 0 {
		t.Error("Tax calculation error, purchase operations do not pay taxes")
	}
	if *outputList[1].Tax > 0 {
		t.Error("Tax calculation error, sales operations with profits below 20000 must not pay taxes")
	}
}

func TestShouldPayTaxesBecauseSellOperationValueBiggerThan20000(t *testing.T) {
	buyOperation := models.NewCapitalGainInput(models.BUY_OPERATION, 10, 100)
	sellOperation := models.NewCapitalGainInput(models.SELL_OPERATION, 350, 100)
	var input []models.CapitalGainInput
	input = append(input, *buyOperation)
	input = append(input, *sellOperation)

	taxCalculation := services.NewTaxCalculation(services.NewBuyOperation(), services.NewSellOperation())

	outputList := taxCalculation.Execute(input)
	if *outputList[0].Tax > 0 {
		t.Error("Tax calculation error, purchase operations do not pay taxes")
	}
	if *outputList[1].Tax == 0 {
		t.Error("Tax calculation error, sales operations with profits above 20000 must pay taxes")
	}
}

func TestShouldNotPayTaxesBecauseSellOperationsWillGenerateLossesAndThenProfitsToCoverTheLosses(t *testing.T) {
	buyOperation := models.NewCapitalGainInput(models.BUY_OPERATION, 10, 10000)
	sellOperation := models.NewCapitalGainInput(models.SELL_OPERATION, 2, 5000)
	var input []models.CapitalGainInput
	input = append(input, *buyOperation)
	input = append(input, *sellOperation)
	sellOperation = models.NewCapitalGainInput(models.SELL_OPERATION, 20, 2000)
	input = append(input, *sellOperation)
	sellOperation = models.NewCapitalGainInput(models.SELL_OPERATION, 20, 2000)
	input = append(input, *sellOperation)

	taxCalculation := services.NewTaxCalculation(services.NewBuyOperation(), services.NewSellOperation())

	outputList := taxCalculation.Execute(input)
	if *outputList[0].Tax > 0 {
		t.Error("Tax calculation error, purchase operations do not pay taxes")
	}
	if *outputList[1].Tax > 0 {
		t.Error("Tax calculation error, sales operations without profits must not pay taxes")
	}
	if *outputList[2].Tax > 0 {
		t.Error("Tax calculation error, sales operations without profits must not pay taxes")
	}
	if *outputList[3].Tax > 0 {
		t.Error("Tax calculation error, sales operations without profits must not pay taxes")
	}
}

func TestShouldReturnTaxError(t *testing.T) {
	buyOperation := models.NewCapitalGainInput(models.BUY_OPERATION, 10, 10000)
	sellOperation := models.NewCapitalGainInput(models.SELL_OPERATION, 20, 11000)
	var input []models.CapitalGainInput
	input = append(input, *buyOperation)
	input = append(input, *sellOperation)

	taxCalculation := services.NewTaxCalculation(services.NewBuyOperation(), services.NewSellOperation())
	output := taxCalculation.Execute(input)

	if output[0].Tax == nil {
		t.Error("nil tax")
	}
	if *output[0].Tax > 0 {
		t.Error("tax should be zero")
	}
	if output[1].Tax != nil {
		t.Error("tax should be nil")
	}
	if output[1].Err == nil {
		t.Error("operation should be returned an error")
	}
	if *output[1].Err == "" {
		t.Error("error should not be empty")
	}
}
