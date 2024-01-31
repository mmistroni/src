package greetings

import (
	"fmt"
	"testing"
)


func TestArea(t *testing.T) {
	
	s := Triangle {base : 5, height : 10}
	
	res2 := s.Area() 
	fmt.Printf("Volume is %f", res2)
	
	if res2  != 25 {
		t.Fatal("Volume not set")
	}

}


func TestChabgeBase(t *testing.T) {
	
	s := Triangle {base : 5, height : 10}
	
	s.changeBase(22)


	res2 := s.Area() 
	fmt.Printf("Volume is %f", res2)
	
	if res2  != 25 {
		t.Fatal("Volume not set")
	}

}

func TestChabgeBasePnt(t *testing.T) {
	
	s := Triangle {base : 5, height : 10}
	
	s.changeBasePnt(20)


	res2 := s.Area() 
	fmt.Printf("Volume is %f", res2)
	
	if res2  != 25 {
		t.Fatal("Volume not set")
	}

}

