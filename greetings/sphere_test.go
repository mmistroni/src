package greetings

import (
	"fmt"
	"testing"
)


func TestVolume(t *testing.T) {
	
	s := Sphere {
		Radius : 5,
	}
	
	res2 := s.Volume() 
	fmt.Printf("Volume is %f", res2)
	
	if res2  == 0 {
		t.Fatal("Volume not set")
	}

}


func TestSurfaceArea(t *testing.T) {
	
	s := Sphere {
		Radius : 5,
	}
	
	res2 := s.Volume() 

	fmt.Printf("Surface is %f", res2)
	if res2 == 0 {
		t.Fatal("Surface area not set")
	}

}




