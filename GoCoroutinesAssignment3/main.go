// GoCoroutinesAssignment3 project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortCoroutine(name string, input []int, channel chan []int) {
	fmt.Printf("%v.About to Sort%v. \n", name, input)
	sort.Ints(input)
	channel <- input
}

func runWithCoroutines(input []int) {
	fmt.Printf("Running with Coroutines for:%v\n", input)
	var holder = []int{}
	sortChannel := make(chan []int)

	partitionSize := len(input) / 4
	fmt.Printf("Array Length: %v, PartitionSize:%v\n", len(input), partitionSize)
	for i := 0; i < 4; i++ {
		start := i * partitionSize
		end := start + partitionSize
		if i == 3 { // last slot, we pick what's remaining
			end = len(input)
		}
		go sortCoroutine("Routine-"+strconv.Itoa(i), input[start:end], sortChannel)
	}

	for i := 0; i < 4; i++ {
		result := <-sortChannel
		holder = append(holder, result...)
	}
	sort.Ints(holder)
	fmt.Printf("Sorted array is:%v", holder)
}

func generateInputArray() []int {
	var numbers = []int{}
	fmt.Print("Enter numbers separated by space character: ")
	if scanner := bufio.NewScanner(os.Stdin); scanner.Scan() {
		// Got the line of text
		var line = scanner.Text()
		if len(line) == 0 {
			fmt.Println("The input is too short!")
			os.Exit(1)
		}

		// Get numbers and store them in the slice
		var words = strings.Split(line, " ")
		for _, word := range words {
			var number, err = strconv.Atoi(word)
			if err != nil {
				fmt.Printf("Failed to parse integer from '%v' and it will be ignored", word)
				continue
			}
			numbers = append(numbers, number)
		}

	} else {
		fmt.Printf("Something went wrong...: %v", scanner.Err())
		os.Exit(1)
	}
	return numbers
}

func main() {
	// Generate random array
	randomArray := generateInputArray()
	fmt.Printf("The random array to sort is:%v\n", randomArray) // at least 10 elements
	runWithCoroutines(randomArray)

}
