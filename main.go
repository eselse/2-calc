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
	var operation string
	for {
		fmt.Print("Enter operation (AVG, SUM, MED): ")
		fmt.Scanln(&operation)
		if operation == "AVG" || operation == "SUM" || operation == "MED" {
			return operation
		} else {
			fmt.Println("Invalid operation. Please enter AVG, SUM, or MED.")
			continue
		}
	}
}

func getSlice() []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter numbers separated by comma: ")

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.Split(input, ",")
		var numbers []int

		for _, p := range parts {
			if n, err := strconv.Atoi(strings.TrimSpace(p)); err == nil {
				numbers = append(numbers, n)
			} else {
				fmt.Println("Invalid number:", p)
			}
		}

		if len(numbers) == 0 {
			fmt.Print("Enter numbers separated by comma: ")
			continue
		}
		return numbers
	}
}

func getResult(operation string, numbers []int) float64 {
	switch operation {
	case "SUM":
		return float64(sum(numbers))
	case "AVG":
		return avg(numbers)
	case "MED":
		return med(numbers)
	default:
		return 0.0
	}
}

func sum(numbers []int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}

func avg(numbers []int) float64 {
	result := 0.0
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	result = float64(sum) / float64(len(numbers))
	return result
}

func med(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0 // or panic, depending on your use case
	}

	// Copy slice to avoid modifying the original
	sorted := make([]int, len(numbers))
	copy(sorted, numbers)
	sort.Ints(sorted)

	n := len(sorted)
	mid := n / 2

	if n%2 == 0 {
		// even length → average of two middle elements
		return float64(sorted[mid-1]+sorted[mid]) / 2.0
	}
	// odd length → middle element
	return float64(sorted[mid])
}
