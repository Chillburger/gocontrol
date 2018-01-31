package main


import (
  "os"
  "strings"
  "io"
  "bytes"
)

var lowercase = []byte("abcdefghijklmnopqrstuvwxyz")
var uppercase = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var lowercase_len = len(lowercase)
var uppercase_len = len(uppercase)


type rot13Reader struct {
  r io.Reader
}

func rot13 (b byte) byte {
  pos := bytes.IndexByte(uppercase, b)

  if pos != -1 {
    return uppercase[(pos+13) % uppercase_len]
  }
  pos = bytes.IndexByte(lowercase, b)
  if pos != -1 {
    return lowercase[(pos+13) % lowercase_len]
  }

  return b
}

func (r rot13Reader) Read(p []byte) (n int, err error) {
  n, err = r.r.Read(p)
  for i :=0 ; i < n; i++ {
    p[i] = rot13(p[i])
  }

  return n, err

}

func main() {
  s := strings.NewReader("Lbh onthoeuoua gur pbqr!")
  r := rot13Reader{s}

  io.Copy(os.Stdout, &r)
}
