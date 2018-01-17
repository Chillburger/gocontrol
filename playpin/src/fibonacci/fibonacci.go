package main

import (
  "fmt"
)


// fibonacci is a function that returns a function
// that returns an int

func fibonacci() func() int {
  counter, prev, current := 0, 0, 1
  next := prev + current
  return func() int {
    var returnValue int
    if counter == 0 {
      counter++
      returnValue = 0
    } else if counter == 1 {
      counter++
      returnValue = 1
    } else {
      returnValue = next
      prev = current
      current = next
      next = prev + current
    }
    return returnValue
  }

}

func main() {
  f := fibonacci()

  for i:= 0; i < 10; i++ {
    fmt.Println(f())
  }
}
