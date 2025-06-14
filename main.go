package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Program for calculating Fibonacci numbers")

	var n int
	strategies := [3]struct {
		name     string
		calcFibo func(int) (uint64, error)
	}{
		{"Recursive", CalcRecursive},
		{"Iterative", CalcIterative},
		{"Binet", CalcBinet},
	}

	for {
		err := getUserNumber(&n)
		if err != nil {
			break
		}

		for _, s := range strategies {
			res, err := s.calcFibo(n)
			fmt.Printf("[%s] Fibonacci(%d) ", s.name, n)
			if err != nil {
				fmt.Printf("could not be calculated: %s\n", err)
				continue
			}
			fmt.Printf("= %d\n", res)
		}
		fmt.Println()
	}
	fmt.Println("Program completed")
}

func getUserNumber(n *int) (err error) {
	fmt.Print("Enter a number [or any to complete]: ")

	var input string
	if _, err = fmt.Scanln(&input); err == nil {
		*n, err = strconv.Atoi(input)
	}
	return
}
