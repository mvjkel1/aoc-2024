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
		panic(e)
	}
}

func readFileContent(fileName string) string {
	data, err := os.ReadFile(fileName)
	check(err)
	return string(data)
}

func parseFileContent() int {
	reExpr := "mul\\(([0-9]+,[0-9]+)\\)"
	re := regexp.MustCompile(reExpr)
	data := readFileContent("input.txt")
	matches := re.FindAllStringSubmatch(data, -1)
	res := 0
	for _, v := range matches {
		nums := strings.Split(v[1], ",")
		n1, err1 := strconv.Atoi(nums[0])
		n2, err2 := strconv.Atoi(nums[1])
		if err1 != nil || err2 != nil {
			log.Fatalf("Invalid input data: %v, %v", err1, err2)
		}
		res += n1 * n2
	}
	return res
}

func part1() {
	fmt.Println(parseFileContent())
}

func main() {
	part1()
}
