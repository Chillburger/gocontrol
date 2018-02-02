package main

// defines the type:
// type Tree struct {
//    Left  *Tree
//    Value int
//    Right *Tree
//  }
import (
  "fmt"
  "golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch
func Walk (t *tree.Tree, ch chan int) {
  traverse(t, ch)
  close(ch)
  // if t.Value != nil {
  //  ch <- t.Value
  //  Walk(t.Left)
  //  Walk(t.Right)
  //}
}

func traverse(t *tree.Tree, ch chan int) {

  if t != nil {
    traverse(t.Left, ch)
    ch <- t.Value
    traverse(t.Right, ch)

  }
}

// Same determines where the trees
// t1 and t2 contain the same values
func Same(t1, t2 *tree.Tree) bool {

  treeOneCh := make(chan int)
  treeTwoCh := make(chan int)

  go Walk(t1, treeOneCh)
  go Walk(t2, treeTwoCh)

  i, y := <-treeOneCh, <-treeTwoCh
}

func main() {
  ch := make(chan int)

  go Walk(tree.New(1), ch)

  for i := range ch {
    fmt.Println("channel values:", i)
  }

  Same(tree.New(1), tree.New(1))
}
