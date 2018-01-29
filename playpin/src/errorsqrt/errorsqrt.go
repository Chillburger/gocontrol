package main

import (
  "fmt"
)

// practice for for loop
func Sqrt(x float64) (float64, error) {
  if x < 0 {
    return 0 , ErrorNegativeSqrt(x)
  }
  var z float64 =  x/10
  for i := 0; i < 10; i ++ {
    z -= (z*z -x) / (2*z);
    fmt.Println("Calculating:", z)
    if z * z == x {
      break
    }
  }
  return z, nil
}


type ErrorNegativeSqrt float64

func (e ErrorNegativeSqrt) Error() string {
  return fmt.Sprint("Couldn't convert number: ", float64(e))
}


func main() {
    if result, err := Sqrt(-2); err != nil {
      fmt.Println(err)
    } else {
      fmt.Println(result)
    }
}
