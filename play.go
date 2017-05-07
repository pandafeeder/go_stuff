//playground and memo for go
package main

import (
	"fmt"
)

//unused global variable will not cause compiler error
var unUsed = 100

func main() {
	play1()
}

func play1() {
	//unused local const will not cause compiler error
	const (
		_ = iota
		one
		two
		three = "III"
		//once iota is interupted, need to recover it explicitly, previous enumeration still counts
		four = iota
		five
	)
	println("five is", five)

	s := "hello世界"
	//byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c,", s[i])
	}
	println()
	//rune
	for _, r := range s {
		fmt.Printf("%c,", r)
	}

	a := [3]int{1, 2, 3}
	//range will copy iterated object
	for i, v := range a {
		if i == 0 {
			a[1], a[2] = 999, 999
			fmt.Println(a)
		}
		//prints out 0,1 | 1,2 | 2, 3
		println(i, v)
	}
	//slice is reference type, doesn't have above issue
	b := []int{1, 2, 3}
	for i, v := range b {
		if i == 0 {
			b[1], b[2] = 999, 999
			fmt.Println(b)
		}
		//prints out 0,1 | 1,999 | 2, 999
		println(i, v)
	}

L1:
	for x := 0; x < 3; x++ {
	L2:
		for y := 0; y < 5; y++ {
			if y > 2 {
				continue L2
			}
			if x > 1 {
				break L1
			}
			println(x, y)
		}
	}
}
