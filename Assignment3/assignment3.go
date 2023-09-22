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
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	sound      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.sound)
}

func (a Animal) DoIt(method string) {
	if method == "eat" {
		a.Eat()
	} else if method == "move" {
		a.Move()
	} else if method == "speak" {
		a.Speak()
	} else {
		fmt.Printf("Invald method called:%v\n", method)
	}
}

func InitAnimals() map[string]Animal {
	holder := map[string]Animal{}
	holder["cow"] = Animal{"grass", "walk", "moo"}
	holder["bird"] = Animal{"worms", "fly", "peep"}
	holder["snake"] = Animal{"mice", "slither", "hss"}
	return holder
}

func ProcessRequest(userInput string, animals map[string]Animal) {
	// split strings to find animal and method
	parts := strings.Split(userInput, " ")

	animal, exist := animals[parts[0]]
	if exist {
		animal.DoIt(parts[1])
	} else {
		fmt.Printf("Invalid Animal:%v\n", animal)
	}

}

func main() {
	fmt.Println("Enter animal and action, separated by space.( Enter Q to finish)")
	scanner := bufio.NewScanner(os.Stdin)
	animalMap := InitAnimals()
	for scanner.Scan() {
		var text = scanner.Text()
		if text == "Q" {
			fmt.Println("Quitting") // Println will add back the final '\n'
			break
		} else {
			ProcessRequest(text, animalMap)

		}

	}

}
