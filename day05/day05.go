package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getOrderingMapping(fileName string) map[int][]int {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", fileName, err)
	}

	mapping := make(map[int][]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		for _, pair := range strings.Fields(scanner.Text()) {
			nums := strings.Split(pair, "|")
			if len(nums) != 2 {
				log.Fatalf("Invalid input format: %s", pair)
			}

			n1, err1 := strconv.Atoi(nums[0])
			n2, err2 := strconv.Atoi(nums[1])
			handleError(err1, "Invalid number format in mapping")
			handleError(err2, "Invalid number format in mapping")

			mapping[n1] = append(mapping[n1], n2)
		}
	}

	return mapping
}

func getCorrectOrderings(mappingFileName string, updatesFileName string) [][]string {
	orderingMapping := getOrderingMapping(mappingFileName)
	file, err := os.Open(updatesFileName)
	defer file.Close()
	handleError(err, fmt.Sprintf("Failed to open file %s", updatesFileName))
	var validOrderings [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, line := range strings.Fields(scanner.Text()) {
			order := strings.Split(line, ",")
			if isOrderingCorrect(order, orderingMapping) {
				validOrderings = append(validOrderings, order)
			}
		}
	}
	return validOrderings
}

func isOrderingCorrect(order []string, mapping map[int][]int) bool {
	for i := 0; i < len(order)-1; i++ {
		n1, err1 := strconv.Atoi(order[i])
		n2, err2 := strconv.Atoi(order[i+1])
		handleError(err1, "Invalid number in ordering")
		handleError(err2, "Invalid number in ordering")
		if !slices.Contains(mapping[n1], n2) {
			return false
		}
	}
	return true
}

func calculateMiddleSum(orderings [][]string) int {
	sum := 0
	for _, path := range orderings {
		mid, err := strconv.Atoi(path[len(path)/2])
		handleError(err, "Invalid number in path")
		sum += mid
	}
	return sum
}

func part1(orderings [][]string) int {
	return calculateMiddleSum(orderings)
}

func getInvalidOrderings(mappingFileName string, updatesFileName string) [][]string {
	orderingMapping := getOrderingMapping(mappingFileName)
	file, err := os.Open(updatesFileName)
	defer file.Close()
	handleError(err, fmt.Sprintf("Failed to open file %s", updatesFileName))
	var inValidOrderings [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, line := range strings.Fields(scanner.Text()) {
			order := strings.Split(line, ",")
			if !isOrderingCorrect(order, orderingMapping) {
				inValidOrderings = append(inValidOrderings, order)
			}
		}
	}
	return inValidOrderings
}

func fixInvalidOrderings(orderingMapping map[int][]int, orderings [][]string) int {
	for _, order := range orderings {
		for !isOrderingCorrect(order, orderingMapping) {
			for i := 0; i < len(order)-1; i++ {
				n1, err1 := strconv.Atoi(order[i])
				n2, err2 := strconv.Atoi(order[i+1])
				handleError(err1, "Invalid number in ordering")
				handleError(err2, "Invalid number in ordering")
				if slices.Contains(orderingMapping[n2], n1) {
					order[i], order[i+1] = order[i+1], order[i]
				}
			}
		}
	}
	return calculateMiddleSum(orderings)
}

func part2(orderings [][]string) int {
	orderingMapping := getOrderingMapping("input1.txt")
	return fixInvalidOrderings(orderingMapping, orderings)
}

func handleError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func main() {
	correctOrderings := getCorrectOrderings("input1.txt", "input2.txt")
	invalidOrderings := getInvalidOrderings("input1.txt", "input2.txt")
	fmt.Println("Part 1 Result:", part1(correctOrderings))
	fmt.Println("Part 2 Result:", part2(invalidOrderings))
}
