// Assignment2 project main.go
package main

import (
	"fmt"
	"math"
	"strconv"
)

func GenDisplaceFn(acceleration, initial_velocity, initial_displacement float64) func() float64 {
	fn := func() float64 {
		// https://courses.fortlewis.edu/courses/17334/pages/kinematics-problem-solving-linear?module_item_id=491717
		/*
					  vf = vi + a t
					  d = vit + ½ a t2
			          vf2 = vi2 + 2 a d

			          vf_sqr := math.Pow(initial_velocity, 2) + 2 * acceleration * initial_displacement
			          vf := math.Sqrt(vf_sqr)
			          time := (vf - initial_velocity) / acceleration
		*/
		vf_sqr := math.Pow(initial_velocity, 2) + 2*acceleration*initial_displacement
		vf := math.Sqrt(vf_sqr)
		return (vf - initial_velocity) / acceleration

	}
	return fn
}

func main() {
	// var then variable name then variable type
	var acceleration, i_velocity, i_displacement string
	// Println function is used to
	// display output in the next line
	fmt.Println("Enter Acceleration: ")
	// Taking input from user
	fmt.Scanln(&acceleration)
	fmt.Println("Enter Initial Velocity: ")
	fmt.Scanln(&i_velocity)
	// Print function is used to
	// display output in the same line
	fmt.Print("Enter Displacement:")
	fmt.Scanln(&i_displacement)

	float_acc, _ := strconv.ParseFloat(acceleration, 32)
	float_ivel, _ := strconv.ParseFloat(i_velocity, 32)
	float_displ, _ := strconv.ParseFloat(i_velocity, 32)

	fun := GenDisplaceFn(float_acc, float_ivel, float_displ)

	fmt.Printf("Time is:%v", fun())

}
