package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const digitsToChoose = 12

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

		maxBank, err := maxBankJoltageK(line, digitsToChoose)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro na linha %q: %v\n", line, err)
			continue
		}

		total += maxBank
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler o arquivo: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(total)
}

func maxBankJoltageK(s string, k int) (int64, error) {
	s = strings.TrimSpace(s)
	n := len(s)
	if n < k {
		return 0, fmt.Errorf("banco com apenas %d dígitos, mas preciso de %d", n, k)
	}

	digits := make([]int, n)
	for i := 0; i < n; i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, fmt.Errorf("caractere inválido %q", s[i])
		}
		digits[i] = int(s[i] - '0')
	}

	toRemove := n - k
	stack := make([]int, 0, n)

	for _, d := range digits {
		for toRemove > 0 && len(stack) > 0 && stack[len(stack)-1] < d {
			stack = stack[:len(stack)-1]
			toRemove--
		}
		stack = append(stack, d)
	}

	if len(stack) > k {
		stack = stack[:k]
	}

	var val int64
	for i := 0; i < k; i++ {
		val = val*10 + int64(stack[i])
	}

	return val, nil
}
