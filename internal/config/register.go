package config

import (
	"capital-gain/internal/services"
)

type Register struct {
	BuyOperation   services.BuyOperation
	SellOperation  services.SellOperation
	TaxCalculation services.TaxCalculation
}

func NewRegister() *Register {
	buyOperation := services.NewBuyOperation()
	sellOperation := services.NewSellOperation()
	taxCalculation := services.NewTaxCalculation(buyOperation, sellOperation)
	return &Register{
		BuyOperation:   buyOperation,
		SellOperation:  sellOperation,
		TaxCalculation: taxCalculation,
	}
}
