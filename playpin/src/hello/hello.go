package main

import (
  "fmt"
  "time"
  "math/rand"
  "math"
)

func add (x int, y int) int {
  return x + y
}

func main() {
  fmt.Printf("hello, world\n")
  fmt.Println("The time is:", time.Now())

  fmt.Println("My Favorite number is: ", rand.Intn(10))
  fmt.Println("test math:", math.Sqrt(7))
  fmt.Println(add(42, 13))

}
