package exercisetest
import (
	"strings"
	"fmt"
)
func Quack(input... string) (string, int) {
	return strings.Join(input, "_"), 1
}

func CallFunction(input func(in string) string, myarg string) string {
	return input(myarg)
}

func Factorial(n int) int {
	if  n == 1 {
		return 1
	}
	return n * Factorial(n - 1)
}

func Looping(start int, end int) int {
	var summer = 0
	for i := start; i <= end; i++ {
		summer += i
	}
	return summer

}	

func ExitWithDefer() int {
	defer fmt.Println("Exiting")
	return 11
}