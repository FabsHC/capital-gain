package main

import (
	"capital-gain/cmd/app/handlers"
	"capital-gain/internal/config"
)

func main() {
	register := config.NewRegister()
	terminalHandler := handlers.NewTerminalHandler(
		register.BuyOperation,
		register.SellOperation,
		register.TaxCalculation)
	terminalHandler.Execute()
}
