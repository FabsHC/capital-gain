package config

import "capital-gain/internal/application/usecase"

type Register struct {
	BuyOperation   usecase.BuyOperation
	SellOperation  usecase.SellOperation
	TaxCalculation usecase.TaxCalculation
}

func NewRegister() *Register {
	buyOperation := usecase.NewBuyOperation()
	sellOperation := usecase.NewSellOperation()
	taxCalculation := usecase.NewTaxCalculation(buyOperation, sellOperation)
	return &Register{
		BuyOperation:   buyOperation,
		SellOperation:  sellOperation,
		TaxCalculation: taxCalculation,
	}
}
