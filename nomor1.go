package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) {
	for x := 1; x <= 10000; x++ {
	for y := 1; y <= 10000-x; y++ {
		for z := 1; z <= 10000/(x*y); z++ {
			if x+y+z == a {
			if x*y*z == b {
			if x*x+y*y+z*z == c {
				fmt.Println(x, y, z)
				return
		    }
		}
	}
	}
}
}
	fmt.Println("no solution")
	return
}

func main() {

	SimpleEquations(1, 2, 3) // no solution

	SimpleEquations(6, 6, 14) // 1 2 3

}
