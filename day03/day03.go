package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func readFileContent(fileName string) string {
	data, err := os.ReadFile(fileName)
	check(err)
	return string(data)
}

func parseMulExpressions(data string, re *regexp.Regexp) int {
	matches := re.FindAllStringSubmatch(data, -1)
	result := 0
	for _, v := range matches {
		nums := strings.Split(v[1], ",")
		n1, err1 := strconv.Atoi(nums[0])
		n2, err2 := strconv.Atoi(nums[1])
		if err1 != nil || err2 != nil {
			log.Fatalf("Invalid input data: %v, %v", err1, err2)
		}
		result += n1 * n2
	}
	return result
}

func part1(data string) int {
	re := regexp.MustCompile(`mul\(([0-9]+,[0-9]+)\)`)
	return parseMulExpressions(data, re)
}

func parseFileContent(data string) string {
	re := regexp.MustCompile(`don't\(\)(.*?)(do\b|$)`)
	matches := re.FindAllStringSubmatch(data, -1)
	for _, match := range matches {
		if len(match) > 1 {
			data = strings.Replace(data, match[1], "", -1)
		}
	}
	return data
}

func part2(data string) int {
	data = parseFileContent(data)
	re := regexp.MustCompile(`mul\(([0-9]+,[0-9]+)\)`)
	return parseMulExpressions(data, re)
}

func main() {
	data := readFileContent("input.txt")
	fmt.Println("Part 1 Result:", part1(data))
	fmt.Println("Part 2 Result:", part2(data))
}
