// Assignment3 project main.go
/*

Write a program which allows the user to get information about a predefined set of animals.
Three animals are predefined, cow, bird, and snake.
Each animal can eat, move, and speak.
The user can issue a request to find out one of three things about an animal:
1) the food that it eats,
2) its method of locomotion, and
3) the sound it makes when it speaks.
The following table contains the three animals and their associated data which should be hard-coded into your program.
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct{}

type Snake struct{}

type Bird struct{}

func (c Cow) Eat() string {
	return "grass"
}

func (c Cow) Move() string {
	return "walk"
}

func (c Cow) Speak() string {
	return "moo"
}

func (s Snake) Eat() string {
	return "mice"
}

func (s Snake) Move() string {
	return "slither"
}

func (s Snake) Speak() string {
	return "hss"
}
func (b Bird) Eat() string {
	return "worms"
}

func (b Bird) Move() string {
	return "fly"
}

func (b Bird) Speak() string {
	return "peep"
}

func DoIt(a Animal, method string) (string, error) {
	switch method {
	case "eat":
		return a.Eat(), nil
	case "move":
		return a.Move(), nil
	case "speak":
		return a.Speak(), nil
	default:
		return "", errors.New("Invalid method called")
	}

}

func Query(parts []string, aninemals map[string]Animal) {
	animal, exist := animals[parts[1]]
	if exist {
		result, err := DoIt(animal, parts[2])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%v", result)
		}
	} else {
		fmt.Printf("Invalid Animal:%v\n", animal)
	}
}

func Create(parts []string) (Animal, error) {
	switch parts[2] {
	case "cow":
		return Cow{}, nil
	case "snake":
		return Snake{}, nil
	case "bird":
		return Bird{}, nil
	default:
		return nil, errors.New("Invalid animal entered")

	}
}

func ProcessRequest(userInput string, animals map[string]Animal) {
	// split strings to find animal and method
	parts := strings.Split(userInput, " ")

	switch parts[0] {
	case "query":
		Query(parts, animals)
	case "newanimal":
		{
			newAnimal, err := Create(parts)
			if err != nil {
				fmt.Println(err)
			} else {
				animals[parts[1]] = newAnimal
				fmt.Println("Created it!")
			}
		}
	default:
		fmt.Printf("Invalid command entered:%v", parts[0])
	}

}

func main() {

	fmt.Println("Enter a command.( Enter Q to finish)")
	fmt.Print(">")
	scanner := bufio.NewScanner(os.Stdin)
	animalMap := map[string]Animal{}
	for scanner.Scan() {
		var text = scanner.Text()
		if text == "Q" {
			fmt.Println("Quitting") // Println will add back the final '\n'
			break
		} else {
			ProcessRequest(text, animalMap)
			fmt.Print("\n>")

		}

	}

}
