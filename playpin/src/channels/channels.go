package main

import (
  "fmt"
)

func sum(s []int, c chan int) {
  sum := 0
  for _, v := range s {
    sum += v
  }
  c <- sum // send sum to c


}

func main() {
  s := []int{7, 2, 8, -9, 4, 0}

  // channels are a typed conduit through which you can send
  // and receive values with the channel operation '<-'
  // channels must be created before use
  c := make(chan int)

  // by default, sends and receives block until the other side is ready
  // allows goroutines to synch
  go sum(s[:len(s)/2], c)
  go sum(s[len(s)/2:], c)
  x, y := <-c, <-c // receive from c
  fmt.Println(x, y, x+y)


  // channels can be buffered, provide the buffer length as the second argument
  ch := make(chan int, 3)

  // sends to a buffered channel block only when the buffer is full
  // receives block when the buffer is empty
  ch <- 1
  ch <- 2
  ch <- 3
  fmt.Println(<-ch)
  fmt.Println(<-ch)
  fmt.Println(<-ch)
}
