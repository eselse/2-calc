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

	// read input line
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// split by comma
	parts := strings.Split(input, ",")

	// convert to slice of ints
	var numbers []int
	for _, p := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err == nil {
			numbers = append(numbers, n)
		} else {
			fmt.Println("Invalid number:", p)
		}
	}
	return numbers
}

func getResult(operation string, numbers []int) float64 {
	result := 0.0
	if operation == "SUM" {
		result = float64(sum(numbers))
	} else if operation == "AVG" {
		result = float64(avg(numbers))
	} else if operation == "MED" {
		result = float64(med(numbers))
	}
	return result
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
