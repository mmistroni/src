package greetings

import (
	"testing"
)


func TestSlices(t *testing.T) {

	
	
	arr := []int{1,2,3,4}
	
	sli := arr[2:]

	sli2 := append(sli, 15)

	if sli2[2] != 15 {
		t.Fatal("Incorrect value for 2ndthelemetn")
	}

}

func TestSlices2(t *testing.T) {

	myslice := make([]int, 4)

	for i:=0; i < len(myslice); i++ {
		myslice[i] = i
	}

	myslice = append(myslice, 44)

	for i:=0; i < len(myslice); i++ {
		if myslice[i] != i {
			t.Fail()
			t.Logf("Incorrect value for %d", i)		
		}
	}	

	
	
	
	
	

}




