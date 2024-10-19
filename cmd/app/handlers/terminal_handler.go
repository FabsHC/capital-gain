package handlers

import (
	"bufio"
	"capital-gain/internal/models"
	"capital-gain/internal/services"
	"encoding/json"
	"os"
)

type TerminalHandler interface {
	Execute()
}

type terminalHandler struct {
	buyOperation   services.BuyOperation
	sellOperation  services.SellOperation
	taxCalculation services.TaxCalculation
}

func NewTerminalHandler(
	buyOperation services.BuyOperation,
	sellOperation services.SellOperation,
	taxCalculation services.TaxCalculation) TerminalHandler {
	return &terminalHandler{
		buyOperation:   buyOperation,
		sellOperation:  sellOperation,
		taxCalculation: taxCalculation,
	}
}

func (t terminalHandler) Execute() {
	scanner := bufio.NewScanner(os.Stdin)
	taxesCalculation := services.NewTaxCalculation(services.NewBuyOperation(), services.NewSellOperation())
	var capitalGainInput []models.CapitalGainInput
	for scanner.Scan() {
		_ = json.Unmarshal([]byte(scanner.Text()), &capitalGainInput)
		capitalGainOutput := taxesCalculation.Execute(capitalGainInput)
		_ = json.NewEncoder(os.Stdout).Encode(capitalGainOutput)
	}
}
