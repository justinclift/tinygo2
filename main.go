package main

import (
	"time"
)

func main () {
	t := time.Now()
	print("Start time: ", t.Hour(), ":", t.Minute(), ":", t.Second(), "\n")
	for i := 0; i < 500000; i++ {
		t = time.Now()
		wasteRam(t)
	}
	t = time.Now()
	print("End time: ", t.Hour(), ":", t.Minute(), ":", t.Second(), "\n")
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
	if f > 10000 {
		println("f has reached 10k")
	}
}



