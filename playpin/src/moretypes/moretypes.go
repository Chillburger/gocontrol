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
}
