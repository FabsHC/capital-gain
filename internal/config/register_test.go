package config_test

import (
	"capital-gain/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRegister(t *testing.T) {
	reg := config.NewRegister()
	assert.NotNil(t, reg)
	assert.NotNil(t, reg.BuyOperation)
	assert.NotNil(t, reg.SellOperation)
	assert.NotNil(t, reg.TaxCalculation)
}
