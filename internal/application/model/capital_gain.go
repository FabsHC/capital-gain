package model

import "ganho-capital/internal/application/util"

type (
	CapitalGainInput struct {
		Operation OperationType `json:"operation"`
		UnitCost  float64       `json:"unit-cost"`
		Quantity  uint          `json:"quantity"`
	}

	CapitalGainOutput struct {
		Tax *float64 `json:"tax,omitempty"`
		Err *string  `json:"err,omitempty"`
	}

	OperationType string
)

const (
	BUY_OPERATION  OperationType = "buy"
	SELL_OPERATION OperationType = "sell"
)

func NewCapitalGainInput(operation OperationType, unitCost float64, quantity uint) *CapitalGainInput {
	return &CapitalGainInput{
		Operation: operation,
		UnitCost:  unitCost,
		Quantity:  quantity,
	}
}

func NewCapitalGainOutput(gains float64) *CapitalGainOutput {
	tax := gains * 0.2
	formattedTax := util.RoundTwoDecimals(tax, 2)
	return &CapitalGainOutput{
		Tax: &formattedTax,
	}
}

func NewCapitalGainOutputError(errMessage string) *CapitalGainOutput {
	return &CapitalGainOutput{
		Err: &errMessage,
	}
}
