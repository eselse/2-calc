package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Count AVG, SUM, MED from slice")
	operation := getOperation()
	numbers := getSlice()
	result := getResult(operation, numbers)
	fmt.Printf("Result: %.2f\n", result)
}

func getOperation() string {
	for {
		fmt.Print("Enter operation (AVG, SUM, MED): ")
		var op string
		fmt.Scanln(&op)

		switch op {
		case "AVG", "SUM", "MED":
			return op
		default:
			fmt.Println("Invalid operation. Please enter AVG, SUM, or MED.")
		}
	}
}

func getSlice() []int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter numbers separated by comma: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if numbers, err := parseNumbers(input); err == nil {
			return numbers
		} else {
			fmt.Println(err)
		}
	}
}

func parseNumbers(input string) ([]int, error) {
	parts := strings.Split(input, ",")
	var numbers []int
	for _, p := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", p)
		}
		numbers = append(numbers, n)
	}
	return numbers, nil
}

func getResult(operation string, numbers []int) float64 {
	switch operation {
	case "SUM":
		return float64(sum(numbers))
	case "AVG":
		return avg(numbers)
	case "MED":
		return med(numbers)
	}
	panic("unsupported operation: " + operation)
}

func sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func avg(numbers []int) float64 {
	return float64(sum(numbers)) / float64(len(numbers))
}

func med(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sorted := append([]int(nil), numbers...) // copy
	sort.Ints(sorted)

	n := len(sorted)
	mid := n / 2

	if n%2 == 0 {
		return float64(sorted[mid-1]+sorted[mid]) / 2.0
	}
	return float64(sorted[mid])
}
