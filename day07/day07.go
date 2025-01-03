package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pair[T, U any] struct {
	First  T
	Second U
}

func convertStringToIntList(s string) []int {
	intList := []int{}
	parts := strings.Split(s, " ")
	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		intList = append(intList, num)
	}
	return intList
}

func readFileContent(fileName string) []Pair[int, []int] {
	file, err := os.Open(fileName)
	defer file.Close()
	handleError(err, fmt.Sprintf("Failed to open file %s", fileName))
	scanner := bufio.NewScanner(file)
	fileContent := []Pair[int, []int]{}
	for scanner.Scan() {
		v := strings.Split(scanner.Text(), ":")
		numComponents := strings.Trim(v[1], " ")
		num, _ := strconv.Atoi(v[0])
		intComponents := convertStringToIntList(numComponents)
		pair := Pair[int, []int]{First: num, Second: intComponents}
		fileContent = append(fileContent, pair)
	}
	return fileContent
}

func handleError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func generateCombinations(numbers []int, index int, currentExpression []string, result *[]string) {
	if index == len(numbers)-1 {
		finalExpression := strings.Join(append(currentExpression, fmt.Sprint(numbers[index])), " ")
		*result = append(*result, finalExpression)
		return
	}
	currentNumber := fmt.Sprint(numbers[index])
	generateCombinations(numbers, index+1, append(currentExpression, currentNumber, "+"), result)
	generateCombinations(numbers, index+1, append(currentExpression, currentNumber, "*"), result)
}

func evalExpr(expr []string) int {
	result, _ := strconv.Atoi(expr[0])
	for i := 1; i < len(expr); i += 2 {
		operator := expr[i]
		num, _ := strconv.Atoi(expr[i+1])
		switch operator {
		case "+":
			result += num
		case "*":
			result *= num
		}
	}

	return result
}
func main() {
	fileContent := readFileContent("input.txt")
	combinations := []string{}
	numToCombinationsPairs := []Pair[int, []string]{}
	for _, pair := range fileContent {
		generateCombinations(pair.Second, 0, []string{}, &combinations)
		numCombinationPair := Pair[int, []string]{First: pair.First, Second: combinations}
		numToCombinationsPairs = append(numToCombinationsPairs, numCombinationPair)
		combinations = []string{}
	}

	results := []int{}
	for _, v := range numToCombinationsPairs {
		for _, combination := range v.Second {
			combinationsList := strings.Split(combination, " ")
			val := evalExpr(combinationsList)
			if val == v.First {
				results = append(results, v.First)
				break
			}
		}
	}
	result := 0
	for _, v := range results {
		result += v
	}
	fmt.Println(result)
}
