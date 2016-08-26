package main

import (
    "fmt"
)
func main() {
    s := []int{0,1,2,3,4,5}
    reverse(&s)
    fmt.Println(s)
}
// reverse reverses an array of ints in place.
// the *[]int is actually a pointer of a reference to slice
func reverse(a *[]int) {
    for i, j := 0, len(*a)-1; i < j; i, j = i+1, j-1{
        // (*a) is a reference to slice
        (*a)[i],(*a)[j] = (*a)[j], (*a)[i]
    }
}