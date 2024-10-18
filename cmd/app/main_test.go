package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ganho-capital/internal/application/model"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestUseCases(t *testing.T) {
	t.Parallel()

	t.Run("should read file resources/case_1 and validate output", func(t *testing.T) {
		stdout := emulateNewTerminalToReadFile("case_1")
		var capitalGainOutput []model.CapitalGainOutput
		if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
			t.Error(err)
		}
	})

	t.Run("should read file resource/case_2 and validate output", func(t *testing.T) {
		stdout := emulateNewTerminalToReadFile("case_2")
		var capitalGainOutput []model.CapitalGainOutput
		if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
			t.Error(err)
		}
	})

	t.Run("should read file resource/case_3 and validate output", func(t *testing.T) {
		stdout := emulateNewTerminalToReadFile("case_3")
		var capitalGainOutput []model.CapitalGainOutput
		if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
			t.Error(err)
		}
	})

	t.Run("should read file resource/case_4 and validate output", func(t *testing.T) {
		stdout := emulateNewTerminalToReadFile("case_4")
		var capitalGainOutput []model.CapitalGainOutput
		if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
			t.Error(err)
		}
	})

	t.Run("should read file resource/case_5 and validate output", func(t *testing.T) {
		stdout := emulateNewTerminalToReadFile("case_5")
		var capitalGainOutput []model.CapitalGainOutput
		if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
			t.Error(err)
		}
	})

	t.Run("should read file resource/case_6 and validate output", func(t *testing.T) {
		stdout := emulateNewTerminalToReadFile("case_6")
		var capitalGainOutput []model.CapitalGainOutput
		if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
			t.Error(err)
		}
	})

	t.Run("should read file resource/case_7 and validate output", func(t *testing.T) {
		stdout := emulateNewTerminalToReadFile("case_7")
		var capitalGainOutput []model.CapitalGainOutput
		if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
			t.Error(err)
		}
	})

	t.Run("should read file resource/case_8 and validate output", func(t *testing.T) {
		stdout := emulateNewTerminalToReadFile("case_8")
		var capitalGainOutput []model.CapitalGainOutput
		if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
			t.Error(err)
		}
	})

	t.Run("should read file resource/case_9 and validate output", func(t *testing.T) {
		stdout := emulateNewTerminalToReadFile("case_9")
		var capitalGainOutput []model.CapitalGainOutput
		if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
			t.Error(err)
		}
	})
}

func emulateNewTerminalToReadFile(fileName string) bytes.Buffer {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Env = os.Environ()

	cmd.Stdin = readFile(fileName)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	return stdout
}

func readFile(fileName string) io.Reader {
	currentPathFile, _ := os.Getwd()
	data, err := os.ReadFile(fmt.Sprintf("%s/../../resources/%s", currentPathFile, fileName))
	if err != nil {
		panic(err)
	}
	return strings.NewReader(string(data))
}
