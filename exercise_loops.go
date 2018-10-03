package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
    z := 1.0
	for i:=0; i<10; i++ {
	    if z<x {
	        z -= (z*z - x) / (2*z)
		}
    	fmt.Println(z)
	}
	return z
}

func Sqrt_v2(x float64) float64 {
    z := 1.0
	for z<x {
	    z -= (z*z - x) / (2*z)
      if z
    	fmt.Println(z)
	}
	return z
}


func main() {
	fmt.Println(Sqrt_v2(2))
	fmt.Println(math.Sqrt(2))
}
