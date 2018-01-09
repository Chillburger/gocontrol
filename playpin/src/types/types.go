package main

import(
  "fmt"
  "math"
  "math/cmplx"
)

// example of a var block
var (
  ToBe bool = false
  MaxInt uint64 = 1<<64 - 1
  z complex128 = cmplx.Sqrt(-5 + 12i)
)

// example of a const block
const(
  Pi = 3.14

  // an untyped constant takes the type needed by context
  // this shifts 1 bit left 100 places to make a huge number
  Big = 1 << 100
  // Shift this big number right 99 places, so it's equivalent to 1<<1
  Small = Big >> 99
)


func needInt(x int) int {
  return x*10 + 1
}

func needFloat(x float64) float64 {
  return x * 0.1
}

func main() {

  // this covers Go types and a few math expressions (bit shifting in particular)
  // shows variable declarations being factored into blocks just like import statements
  fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
  fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
  fmt.Printf("Type: %T Value: %v\n", z, z)

  // This covers var initialization
  var i int
  var f float64
  var b bool
  var s string
  fmt.Printf("%v %v %v %q\n", i, f, b, s)

  // this covers type conversions

  var x, y int = 3, 4
  var floater float64 = math.Sqrt(float64(x*x + y*y))
  var z uint = uint(floater)

  fmt.Println(x, y, z)

  // type inference
  theMagicNumber := -5.555i // change me
  fmt.Printf("The Magic Number is of the type: %T\n", theMagicNumber)

  // constants
  const World ="someword"
  fmt.Println("Hello", World)
  fmt.Println("Happy", Pi, "Day")

  const Truth = true
  fmt.Println("Go is good?", Truth)

  // numeric constants, these are high-precision values
  fmt.Println(needInt(Small))
  fmt.Println(needFloat(Small))
  fmt.Println(needFloat(Big))

}
