package main_test

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
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
		assert.JSONEq(t,
			`[{"tax": 0},{"tax": 0},{"tax": 0}]`,
			string(stdout.Bytes()))
	})

	t.Run("should read file resource/case_2 and validate output", func(t *testing.T) {
		stdout := emulateNewTerminalToReadFile("case_2")
		assert.JSONEq(t,
			`[{"tax": 0.00},{"tax": 10000.00},{"tax": 0.00}]`,
			string(stdout.Bytes()))
	})

	t.Run("should read file resource/case_3 and validate output", func(t *testing.T) {
		stdout := emulateNewTerminalToReadFile("case_3")
		assert.JSONEq(t,
			`[{"tax": 0.00},{"tax": 0.00},{"tax": 1000.00}]`,
			string(stdout.Bytes()))
	})

	t.Run("should read file resource/case_4 and validate output", func(t *testing.T) {
		stdout := emulateNewTerminalToReadFile("case_4")
		assert.JSONEq(t,
			`[{"tax": 0},{"tax": 0},{"tax": 0}]`,
			string(stdout.Bytes()))
	})

	t.Run("should read file resource/case_5 and validate output", func(t *testing.T) {
		stdout := emulateNewTerminalToReadFile("case_5")
		assert.JSONEq(t,
			`[{"tax": 0.00},{"tax": 0.00},{"tax": 0.00},{"tax": 10000.00}]`,
			string(stdout.Bytes()))
	})

	t.Run("should read file resource/case_6 and validate output", func(t *testing.T) {
		stdout := emulateNewTerminalToReadFile("case_6")
		assert.JSONEq(t,
			`[{"tax": 0.00},{"tax": 0.00},{"tax": 0.00},{"tax": 0.00},{"tax": 3000.00}]`,
			string(stdout.Bytes()))
	})

	t.Run("should read file resource/case_7 and validate output", func(t *testing.T) {
		stdout := emulateNewTerminalToReadFile("case_7")
		assert.JSONEq(t,
			`[{"tax":0.00}, {"tax":0.00}, {"tax":0.00}, 
			 {"tax":0.00}, {"tax":3000.00}, {"tax":0.00}, 
			 {"tax":0.00}, {"tax":3700.00}, {"tax":0.00}]`,
			string(stdout.Bytes()))
	})

	t.Run("should read file resource/case_8 and validate output", func(t *testing.T) {
		stdout := emulateNewTerminalToReadFile("case_8")
		assert.JSONEq(t,
			`[{"tax":0.00},{"tax":80000.00},{"tax":0.00},{"tax":60000.00}]`,
			string(stdout.Bytes()))
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
