package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, int) {
	z := 1.0
	i := 20
	for ; i > 0; i-- {
		prev := z
		z -= (z*z - x) / (2 * z)
		if prev == z {
			return z,i
		}
	}
	return z,i
}

func main() {
	x:=64.0
	s,i := Sqrt(x)
	fmt.Println("Sqrt(",x,")=",s,"- Iterations=",i)
}
