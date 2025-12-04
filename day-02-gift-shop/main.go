package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	total := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		ranges := strings.Split(line, ",")

		for _, r := range ranges {
			r = strings.TrimSpace(r)
			if r == "" {
				continue
			}

			parts := strings.Split(r, "-")
			if len(parts) != 2 {
				fmt.Fprintf(os.Stderr, "Intervalo mal formado: %s\n", r)
				continue
			}

			start, err1 := strconv.Atoi(parts[0])
			end, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				fmt.Fprintf(os.Stderr, "Erro ao converter intervalo: %s\n", r)
				continue
			}

			if start > end {
				start, end = end, start
			}

			for i := start; i <= end; i++ {
				if isInvalidID(i) {
					total += i
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler o arquivo: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(total)
}

func isInvalidID(n int) bool {
	s := strconv.Itoa(n)

	if len(s)%2 != 0 {
		return false
	}

	mid := len(s) / 2
	return s[:mid] == s[mid:]
}
