package main

import (
  "fmt"
)


// go functions may be closures
// closure is a function value that references variables from outside its body
// function may access and assign to the referenced values, function is bound to the variables
// the adder function below returns a closure, each closure is bound to its own sum variable
func adder() func(int) int {
    sum := 0
    return func(x int) int {
      sum += x
      return sum
    }
}


func main() {
  pos, neg := adder(), adder()

  for i := 0; i < 10; i++ {
    fmt.Println(pos(i), neg(-2*i))
  }
}
