package handlers

import (
	"bufio"
	"capital-gain/internal/models"
	"capital-gain/internal/services"
	"encoding/json"
	"fmt"
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
	taxCalculation := services.NewTaxCalculation(services.NewBuyOperation(), services.NewSellOperation())

	for scanner.Scan() {
		var capitalGainInputs []models.CapitalGainInput

		err := json.Unmarshal([]byte(scanner.Text()), &capitalGainInputs)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Fail to process input:", err)
			continue
		}

		capitalGainOutputs := taxCalculation.Execute(capitalGainInputs)

		err = json.NewEncoder(os.Stdout).Encode(capitalGainOutputs)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Fail to generate output:", err)
		}
	}

	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Fail to read input:", err)
	}
}
