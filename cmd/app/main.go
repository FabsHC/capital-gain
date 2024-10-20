package main

import (
	"capital-gain/cmd/app/handlers"
	"capital-gain/internal/config"
	"os"
)

func main() {
	register := config.NewRegister()

	terminalHandler := handlers.NewTerminalHandler(
		register.BuyOperation,
		register.SellOperation,
		register.TaxCalculation)

	terminalHandler.Execute(os.Stdin)
}
