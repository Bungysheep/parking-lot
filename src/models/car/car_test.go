package car

import "testing"

func TestCreateCar(t *testing.T) {
	c := New("KA-01-HH-1234", "White")

	if c == nil {
		t.Errorf("Expected to have a new Car.")
	}

	if c.RegNbr != "KA-01-HH-1234" {
		t.Errorf("Expected to have Car Registration Number %s, but got %s", "KA-01-HH-1234", c.RegNbr)
	}

	if c.Colour != "White" {
		t.Errorf("Expected to have Car Colour %s, but got %s", "White", c.Colour)
	}
}

func TestGetCarRegNbr(t *testing.T) {
	c := New("KA-01-HH-1234", "White")

	if c.GetRegNbr() != "KA-01-HH-1234" {
		t.Errorf("Expected to have Car Registration Number %s, but got %s", "KA-01-HH-1234", c.GetRegNbr())
	}
}

func TestGetCarColour(t *testing.T) {
	c := New("KA-01-HH-1234", "White")

	if c.GetColour() != "White" {
		t.Errorf("Expected to have Car Colour %s, but got %s", "White", c.GetColour())
	}
}

func TestCarNotEqual(t *testing.T) {
	c1 := New("KA-01-HH-1234", "White")
	c2 := New("KA-01-BB-0001", "Black")

	if c1.IsEqual(c2) {
		t.Errorf("Expected those cars are not equal, but they are equal.")
	}
}

func TestCarEqual(t *testing.T) {
	c1 := New("KA-01-HH-1234", "White")
	c2 := New("KA-01-HH-1234", "White")

	if !c1.IsEqual(c2) {
		t.Errorf("Expected those cars are equal, but they are not equal.")
	}
}
