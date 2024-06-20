package greetings

import (
	"fmt"
)

func MyFunction(number  int) (int, error) {
	
	if number == -1 {
		panic("Panic attack. interrupting")
	} else if number < 4 {
		err := fmt.Errorf("Something went wrong. %v is below threshold of %v", number, 4)

		return 0, err

	} 
	return number * 2, nil
}