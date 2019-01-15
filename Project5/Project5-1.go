/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import (
  "fmt"
  "time"
)

func main() {
  start := time.Now()
  while := true
  for i := 3 ; while ; i++ {
    b_isEvenlyDivisable := evenlyDivisable(1,20,i)
    if b_isEvenlyDivisable {
      fmt.Println(i)
      fmt.Println(time.Since(start))
      while = false
    }
  }
}

func evenlyDivisable(min,max,num int) bool {
  b_isEvenlyDivisable := true
  for i := min ; i <= max ; i++ {
    if num % i != 0 {
      b_isEvenlyDivisable = false
    }
  }
  return b_isEvenlyDivisable
}
