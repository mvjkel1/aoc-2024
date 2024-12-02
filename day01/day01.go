package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFileContent(fileName string) ([]int, []int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open a file: %v", err)
	}
	defer file.Close()

	var firstColumn, secondColumn []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 2 {
			continue
		}
		f1, err1 := strconv.Atoi(fields[0])
		f2, err2 := strconv.Atoi(fields[1])
		if err1 != nil || err2 != nil {
			log.Fatalf("Invalid input data: %v, %v", err1, err2)
		}
		firstColumn = append(firstColumn, f1)
		secondColumn = append(secondColumn, f2)
	}
	sort.Ints(firstColumn)
	sort.Ints(secondColumn)
	return firstColumn, secondColumn
}

func absDiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func part1(l1, l2 []int) {
	res := 0
	for i := 0; i < len(l1); i++ {
		res += absDiff(l1[i], l2[i])
	}
	fmt.Println("Part 1 Result:", res)
}

func countElement(slice []int, target int) int {
	count := 0
	for _, value := range slice {
		if value == target {
			count++
		}
	}
	return count
}

func part2(l1, l2 []int) {
	res := 0
	for _, val := range l1 {
		res += val * countElement(l2, val)
	}
	fmt.Println("Part 2 Result:", res)
}

func main() {
	l1, l2 := readFileContent("input.txt")
	part1(l1, l2)
	part2(l1, l2)
}
