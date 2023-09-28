package exercisetest_test

import (
	"testing"
	"example.com/golang-exercises/exercisetest"
)

func TestQuack(t *testing.T) {
	testString := "foo"
	if exercisetest.Quack(testString) !=  "FOO" {
		t.Fatal("Wrong Call :(")
	}
}


