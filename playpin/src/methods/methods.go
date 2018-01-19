package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

type MyFloat float64

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

// with the pointer receiver, you change the actual vertex value that is declared in the main method
func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  // go does not have classes but you can define methods on types
  // a method is a function with a special receiver argument
  v := Vertex{ 3, 4}
  fmt.Println(v.Abs())

  // you can declare methods on non struct types Too
  // there is a numeric type MyFloat with an Abs method
  // type must be defined in the same package as the method in order to be a receiver
  // you can't declare a method with a receiver whose type is defined in another package
  f := MyFloat(-math.Sqrt2)
  fmt.Println(f.Abs())

  // you can declare methods with pointer receivers
  // receiver type has the literal syntax *T for some type T
  vert := Vertex { 3, 4}
  vert.Scale(10)
  fmt.Println(vert.Abs())

}
