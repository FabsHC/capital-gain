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

func readFile(t *testing.T, fileName string) io.Reader {
	currentPathFile, _ := os.Getwd()
	data, err := os.ReadFile(fmt.Sprintf("%s/../resources/%s", currentPathFile, fileName))
	if err != nil {
		t.Error(err)
	}
	return strings.NewReader(string(data))
}

// TestCase1 is using a STDIN input with exec.Command to emulate user input with a JSON.
// More details about each case you find in ../docs/CASES.md
func TestCase1(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Env = os.Environ()

	cmd.Stdin = readFile(t, "case_1")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		t.Error(err)
	}

	var capitalGainOutput []model.CapitalGainOutput
	if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
		t.Error(err)
	}

	if *capitalGainOutput[0].Tax > 0 {
		t.Error("Tax calculation error, purchase operations do not pay taxes")
	}
	if *capitalGainOutput[1].Tax > 0 {
		t.Error("Tax calculation error, sales operations with total value below than 20000 must not pay taxes")
	}
	if *capitalGainOutput[2].Tax > 0 {
		t.Error("Tax calculation error, sales operations with total value below than 20000 must not pay taxes")
	}
}

// TestCase2 is using a STDIN input with exec.Command to emulate user input with a JSON.
// More details about each case you find in ../docs/CASES.md
func TestCase2(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Env = os.Environ()

	cmd.Stdin = readFile(t, "case_2")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		t.Error(err)
	}

	var capitalGainOutput []model.CapitalGainOutput
	if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
		t.Error(err)
	}

	if *capitalGainOutput[0].Tax > 0 {
		t.Error("Tax calculation error, purchase operations do not pay taxes")
	}
	if *capitalGainOutput[1].Tax == 0 {
		t.Error("Tax calculation error, sales operations with profits must pay taxes")
	}
	if *capitalGainOutput[2].Tax > 0 {
		t.Error("Tax calculation error, sales operations without profits must not pay taxes")
	}
}

// TestCase3 is using a STDIN input with exec.Command to emulate user input with a JSON.
// More details about each case you find in ../docs/CASES.md
func TestCase3(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Env = os.Environ()

	cmd.Stdin = readFile(t, "case_3")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		t.Error(err)
	}

	var capitalGainOutput []model.CapitalGainOutput
	if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
		t.Error(err)
	}

	if *capitalGainOutput[0].Tax > 0 {
		t.Error("Tax calculation error, purchase operations do not pay taxes")
	}
	if *capitalGainOutput[1].Tax > 0 {
		t.Error("Tax calculation error, sales operations without profits must not pay taxes")
	}
	if *capitalGainOutput[2].Tax == 0 {
		t.Error("Tax calculation error, sales operations with profits must pay taxes")
	}
}

// TestCase4 is using a STDIN input with exec.Command to emulate user input with a JSON.
// More details about each case you find in ../docs/CASES.md
func TestCase4(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Env = os.Environ()

	cmd.Stdin = readFile(t, "case_4")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		t.Error(err)
	}

	var capitalGainOutput []model.CapitalGainOutput
	if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
		t.Error(err)
	}

	if *capitalGainOutput[0].Tax > 0 {
		t.Error("Tax calculation error, purchase operations do not pay taxes")
	}
	if *capitalGainOutput[1].Tax > 0 {
		t.Error("Tax calculation error, purchase operations do not pay taxes")
	}
	if *capitalGainOutput[2].Tax > 0 {
		t.Error("Tax calculation error, sales operations without profits must not pay taxes")
	}
}

// TestCase5 is using a STDIN input with exec.Command to emulate user input with a JSON.
// More details about each case you find in ../docs/CASES.md
func TestCase5(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Env = os.Environ()

	cmd.Stdin = readFile(t, "case_5")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		t.Error(err)
	}

	var capitalGainOutput []model.CapitalGainOutput
	if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
		t.Error(err)
	}

	if *capitalGainOutput[0].Tax > 0 {
		t.Error("Tax calculation error, purchase operations do not pay taxes")
	}
	if *capitalGainOutput[1].Tax > 0 {
		t.Error("Tax calculation error, purchase operations do not pay taxes")
	}
	if *capitalGainOutput[2].Tax > 0 {
		t.Error("Tax calculation error, sales operations without profits must not pay taxes")
	}
	if *capitalGainOutput[3].Tax == 0 {
		t.Error("Tax calculation error, sales operations with profits must pay taxes")
	}
}

// TestCase6 is using a STDIN input with exec.Command to emulate user input with a JSON.
// More details about each case you find in ../docs/CASES.md
func TestCase6(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Env = os.Environ()

	cmd.Stdin = readFile(t, "case_6")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		t.Error(err)
	}

	var capitalGainOutput []model.CapitalGainOutput
	if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
		t.Error(err)
	}

	if *capitalGainOutput[0].Tax > 0 {
		t.Error("Tax calculation error, purchase operations do not pay taxes")
	}
	if *capitalGainOutput[1].Tax > 0 {
		t.Error("Tax calculation error, sales operations without profits must not pay taxes")
	}
	if *capitalGainOutput[2].Tax > 0 {
		t.Error("Tax calculation error, sales operations without profits must not pay taxes")
	}
	if *capitalGainOutput[3].Tax > 0 {
		t.Error("Tax calculation error, sales operations without profits must not pay taxes")
	}
	if *capitalGainOutput[4].Tax == 0 {
		t.Error("Tax calculation error, sales operations with profits must pay taxes")
	}
}

// TestCase7 is using a STDIN input with exec.Command to emulate user input with a JSON.
// More details about each case you find in ../docs/CASES.md
func TestCase7(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Env = os.Environ()

	cmd.Stdin = readFile(t, "case_7")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		t.Error(err)
	}

	var capitalGainOutput []model.CapitalGainOutput
	if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
		t.Error(err)
	}

	if *capitalGainOutput[0].Tax > 0 {
		t.Error("Tax calculation error, purchase operations do not pay taxes")
	}
	if *capitalGainOutput[1].Tax > 0 {
		t.Error("Tax calculation error, sales operations without profits must not pay taxes")
	}
	if *capitalGainOutput[2].Tax > 0 {
		t.Error("Tax calculation error, sales operations without profits must not pay taxes")
	}
	if *capitalGainOutput[3].Tax > 0 {
		t.Error("Tax calculation error, sales operations without profits must not pay taxes")
	}
	if *capitalGainOutput[4].Tax == 0 {
		t.Error("Tax calculation error, sales operations with profits must pay taxes")
	}
	if *capitalGainOutput[5].Tax > 0 {
		t.Error("Tax calculation error, purchase operations do not pay taxes")
	}
	if *capitalGainOutput[6].Tax > 0 {
		t.Error("Tax calculation error, sales operations without profits must not pay taxes")
	}
	if *capitalGainOutput[7].Tax == 0 {
		t.Error("Tax calculation error, sales operations with profits must pay taxes")
	}
	if *capitalGainOutput[8].Tax > 0 {
		t.Error("Tax calculation error, sales operations with total value below than 20000 must not pay taxes")
	}
}

// TestCase8 is using a STDIN input with exec.Command to emulate user input with a JSON.
// More details about each case you find in ../docs/CASES.md
func TestCase8(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Env = os.Environ()

	cmd.Stdin = readFile(t, "case_8")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		t.Error(err)
	}

	var capitalGainOutput []model.CapitalGainOutput
	if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
		t.Error(err)
	}

	if *capitalGainOutput[0].Tax > 0 {
		t.Error("Tax calculation error, purchase operations do not pay taxes")
	}
	if *capitalGainOutput[1].Tax == 0 {
		t.Error("Tax calculation error, sales operations with profits must pay taxes")
	}
	if *capitalGainOutput[2].Tax > 0 {
		t.Error("Tax calculation error, purchase operations do not pay taxes")
	}
	if *capitalGainOutput[3].Tax == 0 {
		t.Error("Tax calculation error, sales operations with profits must pay taxes")
	}
}

func TestCaseInputErr(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Env = os.Environ()

	cmd.Stdin = readFile(t, "inputError")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		t.Error(err)
	}

	var capitalGainOutput []model.CapitalGainOutput
	if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
		t.Error(err)
	}

	if capitalGainOutput[0].Tax == nil {
		t.Error("Tax should not be nil")
	}
	if capitalGainOutput[0].Err != nil {
		t.Error("Error should be nil")
	}
	if *capitalGainOutput[0].Tax > 0 {
		t.Error("Tax calculation error, purchase operations do not pay taxes")
	}
	if capitalGainOutput[1].Tax != nil {
		t.Error("Tax should be nil")
	}
	if capitalGainOutput[1].Err == nil {
		t.Error("Error should not be nil")
	}
	if *capitalGainOutput[1].Err == "" {
		t.Error("Error should not be empty")
	}

}

func TestCaseInputErr2(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Env = os.Environ()

	cmd.Stdin = readFile(t, "inputError")

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		t.Error(err)
	}

	var capitalGainOutput []model.CapitalGainOutput
	if err := json.Unmarshal(stdout.Bytes(), &capitalGainOutput); err != nil {
		t.Error(err)
	}

	if *capitalGainOutput[0].Tax != 0 {
		t.Error("Tax should be zero")
	}
	if *capitalGainOutput[1].Err != "can't sell more stocks than you have" {
		t.Error("Error message invalid")
	}

}
