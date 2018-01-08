package main

import (
  "fmt"
  "time"
  "math/rand"
  "math"
)

var i, j int = 1, 2

func add (x , y int) int {
  return x + y
}

func swap( x, y string) (string, string) {
  return y, x
}

func split ( sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

func main() {
  fmt.Printf("hello, world\n")
  fmt.Println("The time is:", time.Now())

  fmt.Println("My Favorite number is: ", rand.Intn(10))
  fmt.Println("test math:", math.Sqrt(7))
  fmt.Println(add(42, 13))

  fmt.Println("pi:" , math.Pi)
  a, b := swap("hello", "world")
  fmt.Println("swap funct: ", a, b)

  fmt.Println(split(17))
  var c, python, java= true, false, "no!"
  k := 3
  fmt.Println(i, k, j, c, python, java)

}
