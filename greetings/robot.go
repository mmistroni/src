package greetings

import "errors"

type Robot interface {
	PowerOn() error
}

type T850 struct {
	Name string
}

func (a *T850) PowerOn() error {

	return nil
}

type R2D2 struct {
	Broken bool
}

func (a *R2D2) PowerOn() error {

	if a.Broken {
		return errors.New("RD2D is broken")
	} else {
		return nil
	}
}

func Boot(r Robot) error {
	return r.PowerOn()
}


