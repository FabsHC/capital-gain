package services

import (
	"capital-gain/internal/models"
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
	var tax *models.CapitalGainOutput
	purchase := models.NewPurchase(0, 0)
	sale := models.NewSale(0, 0)

	for _, operation := range operations {
		switch operation.Operation {
		case models.BUY_OPERATION:
			tc.buyOperation.Execute(purchase, operation)
			tax = models.NewCapitalGainOutput(0)
		case models.SELL_OPERATION:
			tc.sellOperation.Execute(purchase, sale, operation)
			purchase.SubtractShares(operation.Quantity)
			tax = models.NewCapitalGainOutput(sale.ProfitGains)
		}
		taxes = append(taxes, tax)
		tax = nil
	}
	return taxes
}
