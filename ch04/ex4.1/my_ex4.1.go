package main

import (
    "fmt"
    "crypto/sha256"
)

var pc [256]byte
func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

func PopCount(x uint64) int {
    return int(pc[byte(x>>(0*8))] +
        pc[byte(x>>(1*8))] +
        pc[byte(x>>(2*8))] +
        pc[byte(x>>(3*8))] +
        pc[byte(x>>(4*8))] +
        pc[byte(x>>(5*8))] +
        pc[byte(x>>(6*8))] +
        pc[byte(x>>(7*8))])
}

func main() {
    c1 := sha256.Sum256([]byte("x"))
    c2 := sha256.Sum256([]byte("X"))
    var c3 [32]uint8
    for i := 0; i < len(c1); i++{
        c3[i] = c1[i] ^ c2[i]
    }
    count := 0
    for i := range c3 {
        count += PopCount(uint64(c3[i]))
    }
    fmt.Println(count)
}