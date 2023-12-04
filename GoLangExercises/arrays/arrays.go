package arrays


func Runner() string {
	return "I am running"
}



func CreateArrays(a int) []int {
	var tester =  make([]int, a) 
	tester[0] = 1
	tester[1] = 2
	
	
	return tester
}

func CreateArrays2() []int {
	var tester =  make([]int, 3) 
	tester[0] = 1
	tester[1] = 2
	tester[2] = 3
	
	
	return tester
}

func CreateArraysWithSlice(length int, newelem int) []int {
	var tester =  make([]int, length)
	
	for idx := 0; idx < length; idx++ {
		tester[idx] = idx
	}
	
	tester = append(tester, newelem)
	return tester
}


