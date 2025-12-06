package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	start int
	end   int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao abrir o arquivo: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var intervals []Interval

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}

		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			fmt.Fprintf(os.Stderr, "Intervalo mal formado: %q\n", line)
			os.Exit(1)
		}

		a, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		b, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err1 != nil || err2 != nil {
			fmt.Fprintf(os.Stderr, "Erro ao converter intervalo %q\n", line)
			os.Exit(1)
		}

		if a > b {
			a, b = b, a
		}

		intervals = append(intervals, Interval{start: a, end: b})
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro na leitura: %v\n", err)
		os.Exit(1)
	}

	if len(intervals) == 0 {
		fmt.Println(0)
		return
	}

	merged := mergeIntervals(intervals)
	freshCount := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		id, err := strconv.Atoi(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ID inválido %q\n", line)
			os.Exit(1)
		}

		if isFresh(merged, id) {
			freshCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro na leitura: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(freshCount)
}

func mergeIntervals(intervals []Interval) []Interval {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].start == intervals[j].start {
			return intervals[i].end < intervals[j].end
		}
		return intervals[i].start < intervals[j].start
	})

	var merged []Interval
	for _, cur := range intervals {
		if len(merged) == 0 {
			merged = append(merged, cur)
			continue
		}
		last := &merged[len(merged)-1]

		if cur.start <= last.end+1 {
			if cur.end > last.end {
				last.end = cur.end
			}
		} else {
			merged = append(merged, cur)
		}
	}
	return merged
}

func isFresh(intervals []Interval, id int) bool {
	lo, hi := 0, len(intervals)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		in := intervals[mid]
		if id < in.start {
			hi = mid - 1
		} else if id > in.end {
			lo = mid + 1
		} else {
			return true
		}
	}
	return false
}
