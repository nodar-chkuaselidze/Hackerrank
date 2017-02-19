package main

import "fmt"
import "strings"

func main() {
	var m string
	var e int

	fmt.Scan(&m)
	fmt.Scan(&e)

	fmt.Println(encrypt(m, e))
}

func encrypt(m string, shift int) (c string) {
	A := "0123456789" // Alphabet

	mb := []byte(m)
	cb := make([]byte, len(mb))

	for i, b := range mb {
		cb[i] = A[(strings.IndexByte(A, b)+shift)%len(A)]
	}

	c = string(cb)

	return
}
