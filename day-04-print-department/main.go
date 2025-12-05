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

	var grid []string

	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), "\r\n")
		if strings.TrimSpace(line) == "" {
			continue
		}
		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler o arquivo: %v\n", err)
		os.Exit(1)
	}

	if len(grid) == 0 {
		fmt.Println(0)
		return
	}

	rows := len(grid)
	accessible := 0

	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1} /*{0,0}*/, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for i := 0; i < rows; i++ {
		cols := len(grid[i])
		for j := 0; j < cols; j++ {
			if grid[i][j] != '@' {
				continue
			}

			neighbors := 0

			for _, d := range dirs {
				ni := i + d[0]
				nj := j + d[1]

				if ni < 0 || ni >= rows {
					continue
				}
				if nj < 0 || nj >= len(grid[ni]) {
					continue
				}
				if grid[ni][nj] == '@' {
					neighbors++
				}
			}

			if neighbors < 4 {
				accessible++
			}
		}
	}

	fmt.Println(accessible)
}
