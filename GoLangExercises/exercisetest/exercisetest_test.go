package exercisetest_test

import (
	"testing"
	"example.com/golang-exercises/exercisetest"
	"fmt"
)

func TestQuack(t *testing.T) {
	f := "foo"
	b := "bar"

	str, num := exercisetest.Quack(f, b) 

	if  str !=  "foo_bar" && num !=1 {
		t.Fatal("Wrong Call :(")
	}
}

func TestCallFunction(t *testing.T) {
	var retVal =  "1"
	fn := func(in string)  string{
		return in
	}


	res := exercisetest.CallFunction(fn, retVal)
	if  res !=  retVal {
		t.Fatal("Wrong Call :(")
	}
}

func TestFactorial(t *testing.T) {
	var tester = 5
	res := exercisetest.Factorial(tester)
	if res != 120 {
		fmt.Printf("Expected 120 but was %d", res)
		t.Fail()
	}
}

func TestLooping(t *testing.T) {
	var start = 1
	var end = 10
	res := exercisetest.Looping(start, end)
	if res != 55 {
		fmt.Printf("Expected 55 but was %d", res)
		t.Fail()
	}
}

func TestExitWithDefer(t *testing.T) {
	res := exercisetest.ExitWithDefer()
	if res != 11 {
		fmt.Printf("Expected 55 but was %d", res)
		t.Fail()
	}
}



