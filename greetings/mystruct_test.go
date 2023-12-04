package greetings

import (
	"fmt"
	"testing"
)


func TestCreateMovie(t *testing.T) {
	
	name := "Foobar"
	age := 22

	res2 := CreateMovie(name, age) 

	if res2.Name != name {
		t.Fatal("Wrong Name")
	}	

	newAge := 44

	res2.Age = newAge

	if res2.Age != newAge {
		t.Fatal("Age not set!")
	}

}


func TestReassingt (t *testing.T) {

	var m Movie

    fmt.Printf("%+v\n", m)

    m.Name = "Metropolis"

    m.Age = 11

    fmt.Printf("%+v\n", m)
}


func TestReassing3 (t *testing.T) {

	var m Movie

    fmt.Printf("%+v\n", m)

    m.Name = "Metropolis"

    m.Age = 11

    fmt.Printf("%+v\n", m)
}

func TestCreateMovie2(t *testing.T) {
	
	name := "Foobar"
	age := 23

	res2 := CreateMovie(name, age) 

	if res2.Name != name {
		t.Fatal("Wrong Name")
	}	

	newAge := 44

	res2.Age = newAge

	if res2.Age != newAge {
		t.Fatal("Age not set!")
	}

}

func TestCreateSuperHero(t *testing.T) {
	hero := SuperHero {
		Name : "Batman",
		Age : 22 ,
		Address : Address {
			Street : "foo",
			City : "Gotham",
			Number : 1,
		},

	}

	fmt.Print(hero)
	fmt.Printf("city=%v", hero.Address.City)
}

func TestNewAlarm(t *testing.T) {
	testTime := "TesterTime"
	alarm := NewAlarm(testTime)

	if alarm.Time != testTime {
		t.Fatal("Wrong Name")
	}
	if alarm.Sound != "Klakson" {
		t.Fatal("Wrong Sound")
	}


}

func TestCompareDrinks(t *testing.T ) {

	one := Drink {
		Name : "Lemonade",
		Ice  : true,
	}

	two := Drink {
		Name : "Lemonade",
		Ice  : true,
	}

	three := Drink {
		Name : "Fanta",
		Ice  : true,
	}

	if one != two {
		t.Fatal("Comparison failed. Two items should be same")
	}

	if two == three {
		t.Fatal("Comparison failed. Two items are different")
	}


}

func TestUpdateDrinks(t *testing.T ) {

	one := Drink {
		Name : "Lemonade",
		Ice  : true,
	}

	two := one
	two.Ice = false

	fmt.Printf("One Ice:%v, Two Ice:%v", one.Ice, two.Ice)

	three := &one
	three.Name = "Fanta"


	fmt.Printf("One Name:%v, Two Name:%v, Three Name:%v", one.Name, two.Name, three.Name)


}




