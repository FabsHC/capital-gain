package handlers

import (
	"bufio"
	"capital-gain/internal/models"
	"capital-gain/internal/services"
	"encoding/json"
	"io"
	"os"
)

type TerminalHandler interface {
	Execute(reader io.Reader)
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

func (t terminalHandler) Execute(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	taxCalculation := services.NewTaxCalculation(services.NewBuyOperation(), services.NewSellOperation())

	for scanner.Scan() {
		var capitalGainInputs []models.CapitalGainInput

		_ = json.Unmarshal([]byte(scanner.Text()), &capitalGainInputs)

		capitalGainOutputs := taxCalculation.Execute(capitalGainInputs)

		_ = json.NewEncoder(os.Stdout).Encode(capitalGainOutputs)
	}

}
