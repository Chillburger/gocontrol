package main

import (
  "fmt"
)

func fibonacci(n int, c chan int) {
  x, y := 0, 1
  for i := 0; i < n; i++ {
    c <- x
    x, y = y, x+y
  }
  // a sender can close a channel to indicate that no more values will be sent
  close(c)
}

func main() {
  c := make(chan int, 10)
  go fibonacci(cap(c), c)
  // the loop here receives values from the channel until it is closed
  // checks for v, ok := <-ch where ok is false
  // note that only the sender should close a channel, never the receiver
  // sending on a closed channel will cause a panic
  // channels are not like files, you don't usually need to close them, closing
  // is only necessary when the receiver must be told there ane no more values coming
  for i := range c {
    fmt.Println(i)
  }
}
