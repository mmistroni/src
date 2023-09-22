package main

import (
	"fmt"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}

func main() {
	var a, v0, s0, t float64

	fmt.Print("Acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Init. velocity: ")
	fmt.Scan(&v0)
	fmt.Print("Init. displacement: ")
	fmt.Scan(&s0)

	fn := GenDisplaceFn(a, v0, s0)
	fmt.Println("Created displacement function.")

	fmt.Print("Time:")
	fmt.Scan(&t)
	fmt.Printf("Final displacement: %f\n", fn(t))
}
