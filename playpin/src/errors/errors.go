package main

import (
  "fmt"
  "time"
)


func do (i interface{}) {
  switch v := i.(type) {
  case int:
    fmt.Printf("Twice %v is %v\n", v, v*2)
  case string:
    fmt.Printf("%q is %v bytes long\n", v, len(v))
  default:
    fmt.Printf("I don't know about the type %T\n", v)
  }
}

type Person struct {
  Name string
  Age int
}

func (p Person) String() string {
  return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}


type MyError struct {
  When time.Time
  What string
}

func (e *MyError) Error() string {
  return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
  return &MyError {
    time.Now(),
    "it didn't work",
  }
}

func main() {
  // a type assertion provides access to an interface value's underlying concrete value
  // t := i.(T)
  // if i does not hold a T, the statement will trigger a panic
  var i interface{} = "hello"

  s := i.(string)
  fmt.Println(s)

  s, ok := i.(string)
  fmt.Println(s, ok)

  // a type assertion can return two values when testing where a interface value holds a specific type
  f, ok := i.(float64)

  fmt.Println(f, ok)


  do(21)
  do("hello")
  do(true)

  a := Person{"Aurthur Dent", 42}
  z := Person{"Test Tester", 24}

  fmt.Println(a, z)
  if err := run(); err != nil {
    fmt.Println(err)
  }
}
