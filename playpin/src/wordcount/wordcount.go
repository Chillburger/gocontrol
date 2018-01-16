// word count exercise

package main

import (
  "fmt"
  "golang.org/x/tour/wc"
  "strings"
)

func WordCount (s string) map[string]int {
  wordArray := strings.Fields(s)
  wordCount := make(map[string]int)
  for i := 0; i < len(wordArray); i ++ {
    word, check := wordCount[wordArray[i]]
    if check == true {
      fmt.Println("word found, adding to counter:", wordArray[i], word)
      wordCount[wordArray[i]] = wordCount[wordArray[i]] + 1
    } else {
      fmt.Println("new word found:", wordArray[i], word)
      wordCount[wordArray[i]] = 1
    }
  }

  fmt.Println(wordCount)
  return wordCount
}

func main() {
  wc.Test(WordCount)
}
