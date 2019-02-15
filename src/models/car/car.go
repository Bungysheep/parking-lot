package car

type Car struct {
	RegNbr string
	Colour string
}

// New - Create a new Car
func New(regNbr string, colour string) *Car {
	c := new(Car)
	c.init(regNbr, colour)

	return c
}

func (c *Car) init(regNbr string, colour string) {
	c.RegNbr = regNbr
	c.Colour = colour
}

// GetRegNbr - Return the register number of a car
func (c *Car) GetRegNbr() string {
	return c.RegNbr
}

// GetColour - Return the colour of a car
func (c *Car) GetColour() string {
	return c.Colour
}

// IsEqual - Check whether the other car is equal with the current car
func (c *Car) IsEqual(otherCar *Car) bool {
	return c.RegNbr == otherCar.GetRegNbr() && c.Colour == otherCar.GetColour()
}
