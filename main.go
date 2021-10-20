package main

import (
	"fmt"
	"os"
)

// THIS IS MY FIRST GO PROGRAM
func main() {
	// var s, sep string
	// for i := 0; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
