package main

import ( "code.google.com/p/go-tour/tree"
        "fmt")

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    if t == nil {
        return 
    }
    go Walk(t.Left, ch)
    ch <- t.Value
    go Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool

func main() {
    ch := make(chan int) 
    go Walk(tree.New(10), ch)
    for elem := range ch {
        fmt.Println(elem)
    }
}
