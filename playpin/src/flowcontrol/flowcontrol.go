package main

import "fmt"

func main() {
  // covers go for loops
  sum := 0
  // go for loops don't have pararens surrounding the initial and conditional statements, always need braces around the post statement
  // declare for loop, init statement, conditional statements
  for i := 0; i < 10; i++ {
    // post statement
    sum += i
  }
  fmt.Println(sum)
}
