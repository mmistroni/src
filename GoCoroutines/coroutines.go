// GoCoroutines project main.go
package main

import (
	"fmt"
	"time"
)

func myfunction(label string, count *int) {
	fmt.Printf("%v.Before is %v. \n", label, *count)
	*count += 1
	fmt.Printf("%v.After is %v. \n", label, *count)
}

func main() {
	counter := 0
	go myfunction("First", &counter)
	go myfunction("Second", &counter)

	time.Sleep(5)
	fmt.Println("Counter value = ", counter)

}
