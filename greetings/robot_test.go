package greetings

import (
	"testing"
)


func TestPowerOn(t *testing.T) {
	
	t850 := T850 {Name : "T850"}
	r2d2Good := R2D2{}
	r2d2Bad := R2D2{Broken : true}
	
	
	if t850.PowerOn() != nil {
		t.Fatal("T850 should be nil")
	}

	if r2d2Good.PowerOn() != nil {
		t.Fatal("R2D2 should be nil")
	}


	if r2d2Bad.PowerOn() == nil {
		t.Fatal("R2D2Bad should raise error")
	}


}

func TestBoot(t *testing.T) {

	t850 := T850 {Name : "T850"}
	if Boot(&t850) != nil {
		t.Fatal("T850 should be nil")
	}
}
	


