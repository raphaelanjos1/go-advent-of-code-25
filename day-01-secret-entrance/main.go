package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const dialSize = 100

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao abrir o arquivo: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	current := 50
	passZeroCount := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) < 2 {
			continue
		}

		dir := line[:1]
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Número inválido na linha: %s\n", line)
			continue
		}

		switch dir {
		case "L":
			for i := 1; i <= steps; i++ {
				current = (current - 1 + dialSize) % dialSize
				if current == 0 {
					passZeroCount++
				}
			}
		case "R":
			for i := 1; i <= steps; i++ {
				current = (current + 1) % dialSize
				if current == 0 {
					passZeroCount++
				}
			}
		default:
			fmt.Fprintf(os.Stderr, "Direção inválida na linha: %s\n", line)
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler o arquivo: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(passZeroCount)
}
