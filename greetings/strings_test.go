package greetings

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)



func TestCreateStrings(t *testing.T) {
	
	s := "I am a test string literal"
	
	fmt.Printf("Volume is %s\n", s)


	runeTest := `After a backslash, certain single character escapes represent   
					special values\nn is a line feed or new line \n\t t is a tab`

	fmt.Printf("Rune is %s\n", runeTest)
	

	concatenated := "My name is " + " Slim Shady"
	fmt.Printf("concatenated is %s\n", concatenated)
	
	concatenated += " and appending something else"
	fmt.Printf("concatenated is %s\n", concatenated)
	

}


func TestIntAndStrings(t *testing.T) {
	
	i := 1
	s := "I am a test string literal"
	
	var breakfast string = s + strconv.Itoa(i)


	fmt.Printf("Breakfast is %s\n", breakfast)
}

func TestStringBuffer(t *testing.T) {
	
	var buffer  bytes.Buffer 

	for i:=0;i < 10; i++ {
		buffer.WriteString(strconv.Itoa(i))
	} 



	fmt.Printf("Buffer is %s\n", buffer.String())
}

func TestStringSearch(t * testing.T) {
	fmt.Println(strings.Index("surface", "face"))

	fmt.Println(strings.Index("moon", "aer"))
}


func TestTrimming(t * testing.T) {
	original := "   sur  face   "
	res := strings.TrimSpace(original)
	fmt.Printf("|%s|%s|", original, res)

}




