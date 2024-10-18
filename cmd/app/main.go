package main

import (
	"bufio"
	"encoding/json"
	"ganho-capital/internal/application/model"
	"ganho-capital/internal/application/usecase"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	taxesCalculation := usecase.NewTaxCalculation(usecase.NewBuyOperation(), usecase.NewSellOperation())
	var capitalGainInput []model.CapitalGainInput
	for scanner.Scan() {
		_ = json.Unmarshal([]byte(scanner.Text()), &capitalGainInput)
		capitalGainOutput := taxesCalculation.Execute(capitalGainInput)
		_ = json.NewEncoder(os.Stdout).Encode(capitalGainOutput)
	}
}
