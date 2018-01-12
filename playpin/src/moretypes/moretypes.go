package main

import (
  "fmt"
)

// this is a struct, which is a collection of fields
type Vertex struct {
  X, Y int
}

var (
    // struct literals, denotes a newly allocated struct value by listing vaules of its fields
    v1 = Vertex{1, 2}
    v2 = Vertex{X: 1} // Y:0 is implicit
    v3 = Vertex{} // X:0 and Y: 0
    point = &Vertex{1 ,2} // has type *Vertex
)

func main() {
    // go has pointers!!! here's an example

    i, j := 42, 2701

    p := &i // & generates a pointer to its operand, point to i
    fmt.Println(*p) // read i through the pointer p

    *p = 21 // set i through the pointer
    fmt.Println(i) // new value for i is shown

    p = &j // point to j
    *p= *p / 37 // divide j through the pointer
    fmt.Println(j)

    fmt.Println(Vertex{1 , 2})
    // access a struct field using dot syntax
    v := Vertex{1, 2}
    v.X = 4
    fmt.Println("v.X: ", v.X, " v.Y:", v.Y)


    // pointers to structs
    vert := Vertex{1, 2}
    pointer := &vert
    pointer.X = 1e9
    fmt.Println(vert)

    // from struct literals
    fmt.Println(v1, point, v2, v3)


    // demonstration of arrays
    // this is an array of two strings
    var a [2]string
    a[0] = "hello"
    a[1] = "teamo"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    primes := [6]int{ 2, 3, 5, 7, 11, 13}
    fmt.Println(primes)

    // Slices is a dynamically sized flexible view into elements of an arrays
    // they are much more common in implementation
    var s []int = primes[1:4]

    fmt.Println(s)

    // a slice does not store data, it just describes a section of an underlying arrays
    // changing elements of a silce modifies the corresponding elements of the underlying arrays
    names := [4]string {
      "John",
      "Paul",
      "Ringo",
      "George",
    }

    fmt.Println(names)

    aa := names[0:2]
    bb := names[1:3]

    fmt.Println(aa, bb)

    bb[0] = "XXX"
    fmt.Println(aa, bb)
    fmt.Println(names)

    // slice literals, like an array literal without the length

    q := []int{2, 3, 5, 7, 11, 13}

    fmt.Println(q)

    r := []bool{true, false, true, true, false, true}

    fmt.Println(r)

    ss := []struct {
      i int
      b bool
    }{
      {2, true},
      {3, false},
      {5, true},
      {7, true},
      {11, false},
      {13, true},
    }

    fmt.Println(ss)

    // slice defaults, you can omit the high or low bounds to use their defaults instead

    slicer := []int{2, 3, 5, 7, 11, 13}
    slicer = slicer[1:4]

    fmt.Println(slicer)

    slicer = slicer[:2]
    fmt.Println(slicer)


    slicer = slicer[1:]
    fmt.Println(slicer)

}
