package main

import (
	"time"
)

func main () {
	t := time.Now()
	for i := 0; i < 500000; i++ {
		t = time.Now()
		wasteRam(t)
	}
	t = time.Now()
}

func wasteRam(t time.Time) {
	a := 1
	b := 2
	c := t.Nanosecond()
	d := t.Nanosecond()
	e := (a + c) + (b + d)
	f := 0
	if (e % 7) == 0 {
		f++
	}
}



