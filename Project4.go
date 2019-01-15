/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
  "fmt"
  "strconv"
)

func main() {
  palindromes := findPalindrome(999,999)
  max := findMax(palindromes)
  fmt.Println(max)
}


func findMax(palindromes []int) int {
  max := 0
  for _, v := range palindromes {
    if v > max {
      max = v
    }
  }
  return max
}
func findPalindrome(x, y int) []int {
  var palindromes []int
  for x := 999 ; x > 0 ; x-- {
   for y := 999 ; y > 0 ; y-- {
     if isPalindrome(x*y) {
       palindromes = append(palindromes, x*y)
     }
   }
 }
 return palindromes
}

func isPalindrome(num int) bool {
  sNum := strconv.Itoa(num)
  for i,v := range sNum {
    if string(v) != string(sNum[len(sNum)-(i+1)]) {
      return false
    }
  }
  return true
}
