/*
The sum of the squares of the first ten natural numbers is 385

The square of the sum of the first ten natural numbers is 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

/*
This code found the solution in a matter of micro seconds
*/

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	sumOfSquares := findSumOfSquares(1, 100)
	squareOfSums := findSquareOfSums(1, 100)
	fmt.Println("Sum of Squares: ", sumOfSquares)
	fmt.Println("Square of Sums: ", squareOfSums)
	fmt.Println("Difference: ", sumOfSquares-squareOfSums)
	fmt.Println("Time elapsed: ", time.Since(start))
}

func findSumOfSquares(from, to int) int {
	var sum int
	for i := from; i <= to; i++ {
		sum = sum + int(math.Pow(float64(i), 2))
	}
	return sum
}

func findSquareOfSums(from, to int) int {
	var sum int
	for i := from; i <= to; i++ {
		sum = sum + i
	}
	return int(math.Pow(float64(sum), 2))
}
