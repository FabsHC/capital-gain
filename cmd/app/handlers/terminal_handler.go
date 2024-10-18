package handlers

import (
	"bufio"
	"capital-gain/internal/application/model"
	"capital-gain/internal/application/usecase"
	"encoding/json"
	"os"
)

type TerminalHandler interface {
	Execute()
}

type terminalHandler struct {
	buyOperation   usecase.BuyOperation
	sellOperation  usecase.SellOperation
	taxCalculation usecase.TaxCalculation
}

func NewTerminalHandler(
	buyOperation usecase.BuyOperation,
	sellOperation usecase.SellOperation,
	taxCalculation usecase.TaxCalculation) TerminalHandler {
	return &terminalHandler{
		buyOperation:   buyOperation,
		sellOperation:  sellOperation,
		taxCalculation: taxCalculation,
	}
}

func (t terminalHandler) Execute() {
	scanner := bufio.NewScanner(os.Stdin)
	taxesCalculation := usecase.NewTaxCalculation(usecase.NewBuyOperation(), usecase.NewSellOperation())
	var capitalGainInput []model.CapitalGainInput
	for scanner.Scan() {
		_ = json.Unmarshal([]byte(scanner.Text()), &capitalGainInput)
		capitalGainOutput := taxesCalculation.Execute(capitalGainInput)
		_ = json.NewEncoder(os.Stdout).Encode(capitalGainOutput)
	}
}
