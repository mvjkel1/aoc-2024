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

func readFileContent(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open a file: %v", err)
	}
	defer file.Close()

	var result [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var tmp []int
		for _, v := range strings.Fields(scanner.Text()) {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalln(err)
			}
			tmp = append(tmp, num)
		}
		result = append(result, tmp)
	}
	return result
}

func intsAreSortedDescending(v []int) bool {
	for i := 1; i < len(v); i++ {
		if v[i-1] < v[i] {
			return false
		}
	}
	return true
}

func absDiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func adjacentLevelsDifferOk(v []int) bool {
	for i := 0; i < len(v)-1; i++ {
		diff := absDiff(v[i], v[i+1])
		if !(diff >= 1 && diff <= 3) {
			return false
		}
	}
	return true
}

func part1() int {
	res := 0
	content := readFileContent("input.txt")
	for _, v := range content {
		if sort.IntsAreSorted(v) || intsAreSortedDescending(v) {
			if adjacentLevelsDifferOk(v) {
				res += 1
			}
		}
	}
	return res
}

func main() {
	fmt.Println("Part 1 Result:", part1())
}
