package main

import (
  "fmt"
  "strings"
  "golang.org/x/tour/pic"
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


func printSlice( slice []int) {
  fmt.Printf("len=%d cap=%d %v \n", len(slice), cap(slice), slice)
}

// returns a slice of length dy
// each element of which is a slice of dx 8-bit unsigned integers
func Pic(dx, dy int) [][]uint8 {
    var picture = make([][]uint8, dy)
    for i := 0; i < dy; i++ {
        picture[i] = make([]uint8, dx)
        for j := 0; j < dx; j++ {
          picture[i][j] = uint8(i*j/2)
        }
    }

    return picture
}

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

    // slice length and capacity
    // length of a slice is the number of elements it contains

    sliceExample := []int{2, 3, 5, 7, 11, 13}
    printSlice(sliceExample)

    // slice the slice to give it zero length
    sliceExample = sliceExample[:0]
    printSlice(sliceExample)

    // Extend its length
    sliceExample = sliceExample[:4]
    printSlice(sliceExample)

    // drop firs two values
    sliceExample = sliceExample[1:]
    printSlice(sliceExample)

    sliceExample = sliceExample[:0]
    printSlice(sliceExample)
    sliceExample = sliceExample[:5]
    printSlice(sliceExample)


    // nil slices

    var nilSlice []int
    fmt.Println(nilSlice, len(nilSlice), cap(nilSlice))

    nilSlice = []int{ 1, 2, 3, 4}
    fmt.Println("Setting nilSlice to something:", nilSlice, len(nilSlice), cap(nilSlice))
    if nilSlice == nil {
      fmt.Println("nil!")
    } else {
      fmt.Println("Not nil")
    }

    // creating a slice with make
    // slice can be dynamically created with the built in make FoundationConfiguration
    // this is how you create dynamically sized arrays
    // make function allocates a zeroed array and returns a slice that refers to that array
    aMake := make([]int, 5) // len(aMake)=5
    printSlice(aMake)

    bMake := make([]int, 0, 5) // specifcy a capacity, len(bMake)=0, cap(bMake)=5
    printSlice(bMake)

    cMake := bMake[:2]
    printSlice(cMake)

    dMake := cMake[2:5]
    printSlice(dMake)


    // slices of Slices
    //create a tictac toe board
    board := [][]string {
      []string{"_","_","_"},
      []string{"_","_","_"},
      []string{"_","_","_"},
    }

    // players take turns
    board[0][0] = "X"
    board[2][2] = "O"
    board[1][2] = "X"
    board[1][0] = "O"
    board[0][2] = "X"

    for i := 0; i < len(board); i++ {
      fmt.Printf("%s\n", strings.Join(board[i], " "))
    }


    // appending to a sliceExample
    // it's common to add new elements to a slice, using the append FoundationConfiguration

    var appendSlice []int
    printSlice(appendSlice)

    appendSlice = append(appendSlice, 0)
    printSlice(appendSlice)

    appendSlice = append(appendSlice, 1)
    printSlice(appendSlice)

    appendSlice = append(appendSlice, 1, 2, 3, 4)
    printSlice(appendSlice)

    // range form of the for loop interates over a slice or a map

    var pow  = []int{1, 2, 4, 8, 16, 32, 64, 128}

    for i, v := range pow {
      fmt.Printf("2**%d = %d\n", i, v)
    }

    // you can skip the index or value by assigning to _

    powDrop := make([]int, 10)

    for i := range powDrop {
      powDrop[i] = 1 << uint(i) // 2**i
    }
    for _, value := range powDrop {
      fmt.Printf("%d\n", value)
    }


    // Slice exercise
    // implement Pic, should return a slice of length dy, each element of which
    // is a slice of dx 7-bit unsigned integers
    // When you run the program, it will display your picture, interpretting the integers
    // as grayscale values

    // same functions: (x+y)/2, x*y, x^y
    // need to use a loop to allocate each []uint8 instide the [][]uint8
    // use uint8(intValue) to convert between types
    pic.Show(Pic)

}
