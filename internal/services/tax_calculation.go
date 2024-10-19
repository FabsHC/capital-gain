package services

import (
	"capital-gain/internal/models"
	"capital-gain/internal/utils"
)

type TaxCalculation interface {
	Execute(operations []models.CapitalGainInput) []*models.CapitalGainOutput
}

type (
	taxCalculation struct {
		buyOperation  BuyOperation
		sellOperation SellOperation
	}
)

func NewTaxCalculation(buyOperation BuyOperation, sellOperation SellOperation) TaxCalculation {
	return &taxCalculation{
		buyOperation:  buyOperation,
		sellOperation: sellOperation,
	}
}

func (tc *taxCalculation) Execute(operations []models.CapitalGainInput) []*models.CapitalGainOutput {
	var taxes []*models.CapitalGainOutput
	purchase := models.NewPurchase()
	sale := models.NewSale()

	for _, operation := range operations {
		var tax *models.CapitalGainOutput

		switch operation.Operation {
		case models.BUY_OPERATION:
			tc.buyOperation.Execute(purchase, operation)
			tax = models.NewCapitalGainOutput(utils.ZERO)

		case models.SELL_OPERATION:
			tc.sellOperation.Execute(purchase, sale, operation)
			purchase.SubtractShares(operation.Quantity)
			tax = models.NewCapitalGainOutput(getSalesGains(sale, operation))
		}

		taxes = append(taxes, tax)
	}

	return taxes
}

func getSalesGains(sale *models.Sale, operation models.CapitalGainInput) float64 {
	if operation.GetTotalCost(operation.UnitCost) <= 20000 {
		return utils.ZERO
	}

	return sale.Gains
}
