package main

import (
	"os"
)

func main () {
	n, err := os.OpenFile("/dev/null", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer n.Close()
	for i := 0; i < 100000; i++ {
		foo(n)
	}
}

func foo(n *os.File) {
	a := "1"
	b := "2"
	c := a + b
	numBytes, err := n.WriteString(c)
	if err != nil {
		panic(err)
	}
	if numBytes != 2 {
		panic("Incorrect number of bytes written")
	}
}



