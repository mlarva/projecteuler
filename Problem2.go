package main

import (
  "fmt"
)

func main() {
  sequence := CreateFibonacciSequence(4000000)
  fmt.Printf("\n%v", sequence)
  sum := SumIntArray(sequence)
  fmt.Println("Sum:", sum)
}

func CreateFibonacciSequence(upto int) []int {
  var sequence []int
  var sum int
  lessthan := true
  sequence = append(sequence, 1)
  sequence = append(sequence, 2)
  for i := 0 ; lessthan ; i++ {
    sum = sequence[i] + sequence[i+1]
    if sum > 4000000 {
      lessthan = false
      return sequence
    } else {
      sequence = append(sequence, sum)
    }
  }
  return sequence
}

func SumIntArray(array []int) int {
  var sum int
  for _, v := range array {
    if v % 2 == 0 {
          sum = sum + v
    }
  }
  return sum
}
