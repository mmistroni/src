package greetings

type Movie struct {
	Name string
	Age  int
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