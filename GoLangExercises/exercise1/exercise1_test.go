package exercise1_test
import (
	"example.com/golang-exercises/exercise1"
	"fmt"
	"testing"
	"reflect"
	
	
)
func TestMascot(t *testing.T) {
	if exercise1.Runner() ==  "Go Gopher" {
		t.Fatal("Wrong Mascot :(")
	}
}

func TestType(t *testing.T) {
	var ig int = 32
	fmt.Println(reflect.TypeOf(ig))

}

func TestConvert(t *testing.T) {
	var ig int = 32

	
	fmt.Println(reflect.TypeOf(ig))
	
}