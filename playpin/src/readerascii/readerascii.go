package main

import (
  "golang.org/x/tour/reader"
)

type MyReader struct{}




// add a Read([]byte) (int, error) method to MyReader

func (read MyReader) Read(s []byte) ( int, error) {
  for i := 0; i < len(s); i++ {
    s[i] = 'A'
  }
  return len(s), nil


}

func main() {
  reader.Validate(MyReader{})
}
