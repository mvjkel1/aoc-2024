package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFileContent(fileName string) [][]string {
	file, err := os.Open(fileName)
	defer file.Close()
	handleError(err, fmt.Sprintf("Failed to open file %s", fileName))
	scanner := bufio.NewScanner(file)
	var routeMap [][]string
	for scanner.Scan() {
		for _, line := range strings.Fields(scanner.Text()) {
			routeMap = append(routeMap, strings.Split(line, ""))
		}
	}
	return routeMap
}

func getGuardPosAndDir(routeMap [][]string) (string, int, int) {
	for y := 0; y < len(routeMap); y++ {
		for x := 0; x < len(routeMap[y]); x++ {
			if routeMap[y][x] == "^" {
				return "north", x, y
			}
		}
	}
	return "other planet", -1, -1
}

func changeGuardDir(guardDir string) string {
	switch guardDir {
	case "north":
		return "east"
	case "east":
		return "south"
	case "west":
		return "north"
	case "south":
		return "west"
	default:
		return "other planet"
	}
}

func contains(slice [][]int, target []int) bool {
	for _, elem := range slice {
		if elem[0] == target[0] && elem[1] == target[1] {
			return true
		}
	}
	return false
}

func walk(routeMap [][]string) {
	sum := 0
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Part 1:", sum)
		}
	}()
	var visitedPosMap [][]int
	direction, guardX, guardY := getGuardPosAndDir(routeMap)
	move := map[string]struct{ dx, dy int }{
		"north": {0, -1},
		"east":  {1, 0},
		"west":  {-1, 0},
		"south": {0, 1},
	}
	for {
		step := move[direction]
		for nx, ny := guardX+step.dx, guardY+step.dy; routeMap[ny][nx] != "#"; nx, ny = nx+step.dx, ny+step.dy {
			guardX, guardY = nx, ny
			if !contains(visitedPosMap, []int{guardX, guardY}) {
				sum++
				visitedPosMap = append(visitedPosMap, []int{guardX, guardY})
			}
		}
		direction = changeGuardDir(direction)
	}
}

func part1() {
	routeMap := readFileContent("input.txt")
	walk(routeMap)
}

func handleError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}

func main() {
	part1()
}
