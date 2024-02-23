package usecase

import (
	"ganho-capital/internal/application/model"
)

type (
	TaxCalculation struct {
		buyOperation  *BuyOperation
		sellOperation *SellOperation
	}
)

func NewTaxCalculation(buyOperation *BuyOperation, sellOperation *SellOperation) *TaxCalculation {
	return &TaxCalculation{
		buyOperation:  buyOperation,
		sellOperation: sellOperation,
	}
}

func (tc *TaxCalculation) Execute(operations []model.CapitalGainInput) []*model.CapitalGainOutput {
	var taxes []*model.CapitalGainOutput
	var tax *model.CapitalGainOutput
	purchase := model.NewPurchase(0, 0)
	sale := model.NewSale(0, 0)

	for _, operation := range operations {
		switch operation.Operation {
		case model.BUY_OPERATION:
			tc.buyOperation.Execute(purchase, operation)
			tax = model.NewCapitalGainOutput(0)
			break
		case model.SELL_OPERATION:
			if saleErr := sale.Validate(purchase.TotalShares, operation.Quantity); saleErr != nil {
				tax = model.NewCapitalGainOutputError(saleErr.Error())
				break
			}
			tc.sellOperation.Execute(purchase, sale, operation)
			purchase.SubtractShares(operation.Quantity)
			tax = model.NewCapitalGainOutput(sale.ProfitGains)
			break
		}
		taxes = append(taxes, tax)
		tax = nil
	}
	return taxes
}
