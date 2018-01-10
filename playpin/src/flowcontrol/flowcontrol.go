package main

import (
  "fmt"
  "math"
  "runtime"
  "time"
)

// demonstrates an if statement, braces are required, the expression does not need pararens
func  sqrt(x float64) string  {
  if x < 0 {
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  }
  return lim
}


func powElse(x, n, lim float64) float64 {
  // else statements have access to variables declared in an if statement
  // both calls to powElse are executed and returned before the call to fmt.Println in main begins
  if v := math.Pow(x, n); v < lim {
    return v
  } else {
    fmt.Printf( "%g >= %g\n", v, lim)
  }

  return lim
}

// practice for for loop
func sqrtPractive(x float64) float64 {
  var z float64 =  x/10
  for i := 0; i < 10; i ++ {
    z -= (z*z -x) / (2*z);
    fmt.Println("Calculating:", z)
    if z * z == x {
      break
    }
  }
  return z
}

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


  // demonstrate that init/post statements are optional
  sumoptional := 1
  for; sumoptional < 1000; {
    sumoptional += sumoptional

  }
  fmt.Println("sum:", sumoptional)


  // demonstrate that For is Go's 'while' loop equivalent

  sumwhile := 1
  for sumwhile < 1000 {
    sumwhile += sumwhile
  }
  fmt.Println("Somewhile:", sumwhile)




  // if statements
  fmt.Println(sqrt(2), sqrt(-4))

  // if with a short statement
  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
  )

  // if / else statements
  fmt.Println(
    powElse(3, 2, 10),
    powElse(3, 3, 20),
  )

  // for loop practice
  fmt.Println(sqrtPractive(4))

  // switch statements
  fmt.Print("Go runs on: ")
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OSX")
  case "linux":
    fmt.Println("Linux")
  default:
    fmt.Printf("%s\n", os)
  }


  // switch evaluation order
  fmt.Println("When's Saturday?")
  today := time.Now().Weekday()
  switch time.Saturday {
  case today + 0:
    fmt.Println("Today")
  case today + 1:
    fmt.Println("Tomorrow")
  case today + 2:
    fmt.Println("Two Days")
  default:
    fmt.Println("Too far away.")
  }

  // swtich statement with no conditional
  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("Good morning")
  case t.Hour() < 17:
    fmt.Println("Good afternoon")
  default:
    fmt.Println("Good evening")
  }

  // switch statement defer

  defer fmt.Println("world")

  fmt.Println("hello")

  // stacking defers
  fmt.Println("counting")

  for counter := 0; counter < 10; counter++ {
    defer fmt.Println(counter)
  }
  fmt.Println("done")
}
