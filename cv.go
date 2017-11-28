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

func children(m map[int][]string) {
	for i := 1; i <= 25; i++ {
		children := make([]int, 0)
		fmt.Printf("ID: %v\n", i)
		for _, c := range m[i] {
			numChild := stringToInt(c)
			children = append(children, numChild)
			for _, g := range m[numChild] {
				numGrand := stringToInt(g)
				children = append(children, numGrand)
			}
		}
		sort.Ints(children)
		fmt.Printf("Children: %v\n", children)
	}
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := make(map[int][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), " ")
		i := stringToInt(nums[0])
		m[i] = nums[1:]
	}
	children(m)
}
