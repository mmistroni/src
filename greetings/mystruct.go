package greetings

import "strconv"

type Movie struct {
	Name   string
	Age    int
	Rating float64
}

type SuperHero struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Number int
	Street string
	City   string
}

type Alarm struct {
	Time  string
	Sound string
}

func NewAlarm(time string) Alarm {
	return Alarm{
		Time:  time,
		Sound: "Klakson",
	}
}

type Drink struct {
	Name string
	Ice  bool
}

func CreateMovie(name string, age int) Movie {
	m := Movie{
		Name: name,
		Age:  age,
	}
	return m
}

func CreateMovie2(name string, age int) *Movie {
	m := new(Movie)
	m.Name = name
	m.Age = age
	return m
}

func (m *Movie) summary() string {

	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)

	return m.Name + ", " + r

}