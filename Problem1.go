package main

import (
  "fmt"
)

func main() {
  result := FindMultiplesOfThreeOrFive(1000)
  sum := ComputeSum(result)
  fmt.Println("Sum:", sum)
}

func FindMultiplesOfThreeOrFive(upto int) []int {
  var result []int
  for i:= 1 ; i < upto ; i++ {
    if i % 3 == 0 || i% 5 == 0 {
      result = append(result, i)
    }
  }
  return result
}



func ComputeSum(result []int) int {
  var sum int
  for _, v := range result {
    sum = sum + v
  }
  return sum
}
