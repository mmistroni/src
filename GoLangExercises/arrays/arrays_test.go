package arrays

import (
	"testing"
	"example.com/arrays"
	
)

func TestCreateArray(t *testing.T) {
	
	testLength := 2

	res2 := arrays.CreateArrays(testLength)

	if len(res2) != testLength {
		t.Fatal("Wrong Length")
	}	
}

func TestCreateArrayCopy(t *testing.T) {
	
	testLength := 2

	res2 := arrays.CreateArrays(testLength)

	if len(res2) != testLength {
		t.Fatal("Wrong Length")
	}	
}


func TestCreateArray2(t *testing.T) {
	
	testLength := 3

	res2 := arrays.CreateArrays2()

	if len(res2) != testLength {
		t.Fatal("Wrong Length")
	}	
}

func TestCreateArrayWithSlice(t *testing.T) {
	
	testLength := 2

	testElem := 99

	res2 := arrays.CreateArraysWithSlice(testLength, testElem)

	res2 = append(res2, 1,2,3,4,5)




	if len(res2) != testLength + 6 {
		t.Fatal("Wrong Length")
	}	

	if res2[testLength] != testElem {
		t.Fatal("testelement not found")
	}

}

func TestRemoveFromArray(t *testing.T) {
	
	t.Log("sample")
	testLength := 2

	testElem := 99

	res2 := arrays.CreateArraysWithSlice(testLength, testElem)

	t.Log("Before Removing ")

	res2 = append(res2[:1], res2[2:]...)

	t.Log(res2)

	if len(res2) != testLength {
		t.Fatal("Wrong Length")
	}





}

