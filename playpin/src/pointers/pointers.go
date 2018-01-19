package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

func Abs(v Vertex) float64 {
  return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// rewritten as function
func Scale (v *Vertex, f float64) {
  v.X = v.X * X
  v.Y = v.Y * Y
}


func main() {

  v := Vertex {3, 4}
  // methods with a pointer arguement must take a pointer
  Scale(&v, 10)
  fmt.Println(Abs(v))

}
