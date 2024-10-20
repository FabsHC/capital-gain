package handlers_test

import (
	"capital-gain/cmd/app/handlers"
	"capital-gain/internal/config"
	"capital-gain/internal/models"
	"encoding/json"
	"errors"
	"github.com/agiledragon/gomonkey/v2"
	_ "github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strings"
	"testing"
)

var (
	validInput   = "[{\"operation\":\"buy\", \"unit-cost\":10.00, \"quantity\": 100},{\"operation\":\"sell\", \"unit-cost\":10.00, \"quantity\": 100}]"
	invalidInput = "[{]"
)

func TestHandler(t *testing.T) {
	t.Parallel()
	reg := config.NewRegister()

	t.Run("should execute with success with case_1 data", func(t *testing.T) {
		handler := handlers.NewTerminalHandler(reg.BuyOperation, reg.SellOperation, reg.TaxCalculation)
		var inputReader io.Reader = strings.NewReader(validInput)

		reader, writer, err := os.Pipe()
		assert.Nil(t, err)
		originalStdout := os.Stdout
		os.Stdout = writer

		handler.Execute(inputReader)
		err = writer.Close()
		assert.Nil(t, err)

		out, err := io.ReadAll(reader)
		assert.Nil(t, err)

		var capitalGainOutputs []models.CapitalGainOutput
		err = json.Unmarshal(out, &capitalGainOutputs)
		assert.Nil(t, err)

		assert.Equal(t, 2, len(capitalGainOutputs))
		assert.Equal(t, 0.0, *capitalGainOutputs[0].Tax)
		assert.Equal(t, 0.0, *capitalGainOutputs[1].Tax)

		os.Stdout = originalStdout
	})

	t.Run("should execute with failure with invalid validInput data", func(t *testing.T) {
		handler := handlers.NewTerminalHandler(reg.BuyOperation, reg.SellOperation, reg.TaxCalculation)

		var inputReader io.Reader = strings.NewReader(invalidInput)

		reader, writer, err := os.Pipe()
		assert.Nil(t, err)
		originalStdout := os.Stdout
		os.Stdout = writer

		handler.Execute(inputReader)
		err = writer.Close()
		assert.Nil(t, err)

		out, err := io.ReadAll(reader)
		assert.Nil(t, err)

		assert.Empty(t, out)

		os.Stdout = originalStdout
	})

	t.Run("should execute with failure, json encode error", func(t *testing.T) {
		gomonkey.ApplyFunc((*json.Encoder).Encode, func(_ interface{}) error {
			return errors.New("fail")
		})

		handler := handlers.NewTerminalHandler(reg.BuyOperation, reg.SellOperation, reg.TaxCalculation)

		var inputReader io.Reader = strings.NewReader(validInput)

		reader, writer, err := os.Pipe()
		assert.Nil(t, err)
		originalStdout := os.Stdout
		os.Stdout = writer

		handler.Execute(inputReader)
		err = writer.Close()
		assert.Nil(t, err)

		out, err := io.ReadAll(reader)
		assert.Nil(t, err)

		assert.Empty(t, out)

		os.Stdout = originalStdout
	})
}
