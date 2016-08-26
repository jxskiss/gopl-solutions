package main

import (
    "fmt"
)
func main() {
    //note that byte has 8 bit.
    s := []byte{0,1,2,3,4,5,255}
    reverse(s)
    fmt.Println(s)
}
// reverse reverses an array of bytes in place.
func reverse(a []byte) {
    for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1{
        a[i],a[j] = a[j], a[i]
    }
}