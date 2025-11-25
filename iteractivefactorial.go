package main

import "fmt"

func iterativeFactorial(n int) int {
	result := 1

	for i := 2; i <= n; i++ {
		result = result * i
	}
	return  result
}

func factorialRecursive(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return  n * factorialRecursive(n-1)
}

func main() {
	n := 5
	fmt.Println("Iterative:", iterativeFactorial(n))
	fmt.Println("Recursive:", factorialRecursive(n))
}
