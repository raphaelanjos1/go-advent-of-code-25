package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao abrir o arquivo: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var total int64

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		maxBank, err := maxBankJoltage(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro na linha %q: %v\n", line, err)
			continue
		}

		total += int64(maxBank)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler o arquivo: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(total)
}

func maxBankJoltage(s string) (int, error) {
	s = strings.TrimSpace(s)
	if len(s) < 2 {
		return 0, fmt.Errorf("banco com menos de 2 baterias")
	}

	n := len(s)
	digits := make([]int, n)

	for i := 0; i < n; i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, fmt.Errorf("caractere inválido %q", s[i])
		}
		digits[i] = int(s[i] - '0')
	}

	sufMax := make([]int, n)
	sufMax[n-1] = digits[n-1]
	for i := n - 2; i >= 0; i-- {
		if digits[i] > sufMax[i+1] {
			sufMax[i] = digits[i]
		} else {
			sufMax[i] = sufMax[i+1]
		}
	}

	maxVal := -1
	for i := 0; i < n-1; i++ {
		tens := digits[i]
		units := sufMax[i+1]
		val := tens*10 + units
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal, nil
}
