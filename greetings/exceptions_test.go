package greetings

import (
	"fmt"
	"testing"
)



func TestCallFunctionNoException(t *testing.T) {
	
	good := 5
	res, err := MyFunction(5)

	if err  != nil {
		t.Fail()
		t.Logf("there should not be any errors for calling function with  %d", good)		
	} else {
		fmt.Print(res)
	}
}

func TestCallFunctionWithException(t *testing.T) {
	
	good := 3
	res, err := MyFunction(good)

	if err  == nil {
		t.Fail()
		t.Logf("there should not be any errors for calling function with  %d", good)		
	} else {
		fmt.Println(err)
		fmt.Println(res)
	}
}

func TestCallFunctionWithPanic(t *testing.T) {
	
	good := -1
	res, err := MyFunction(good)

	if err  == nil {
		t.Fail()
		t.Logf("there should not be any errors for calling function with  %d", good)		
	} else {
		fmt.Println(err)
		fmt.Println(res)
	}
}



