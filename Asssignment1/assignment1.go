// Asssignment1 project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Swap(slice []int, index int) {
	tmp := slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = tmp
}

func BubbleSort(myslice []int) {

	loop := len(myslice)

	var swapped = false

	for i := 0; i < loop-1; i++ {
		swapped = false

		for j := 0; j < loop-1; j++ {

			/* compare the adjacent elements */
			if myslice[j] > myslice[j+1] {
				/* swap them */
				Swap(myslice, j)
				swapped = true
			}

		}

		if !swapped {
			break
		}

	}

}

func main() {
	fmt.Println("Enter up to 10 numbers( Enter Q to finish)")
	scanner := bufio.NewScanner(os.Stdin)
	holder := []int{}
	idx := 0
	for scanner.Scan() {
		if idx == 10 {
			fmt.Println("Max length reached")
			break
		}
		var text = scanner.Text()
		if text == "Q" {
			fmt.Println("Quitting") // Println will add back the final '\n'
			break
		} else {
			value, _ := strconv.Atoi(text)
			fmt.Println(value)
			holder = append(holder, value)
			idx += 1
		}

	}
	fmt.Printf("Starting calculation for:%v", holder)
	BubbleSort(holder)

	fmt.Printf("And the result is:%v", holder)

}
