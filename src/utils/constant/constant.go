package constant

const (
	// EmptyString - Blank string
	EmptyString = ""

	// WhiteSpace - White space
	WhiteSpace = " "

	// NewLine - New line
	NewLine = "\n"

	// Comma - Comma
	Comma = ", "

	// MinimumCapacity - Minimum Capacity of a parking lot
	MinimumCapacity = 1

	// CreateParkingLotAction - Action to create a parking lot
	CreateParkingLotAction = "create_parking_lot"

	// ParkAction - Action to park a car into parking lot
	ParkAction = "park"

	// LeaveAction - Action to leave a car from parking lot
	LeaveAction = "leave"

	// StatusAction - Action to return status of a parking lot
	StatusAction = "status"

	// RegistrationNumberOfCarsByColourAction - Action to return registration number of cars by colour
	RegistrationNumberOfCarsByColourAction = "registration_numbers_for_cars_with_colour"

	// SlotNumberOfCarsByColourAction - Action to return slot number of cars by colour
	SlotNumberOfCarsByColourAction = "slot_numbers_for_cars_with_colour"

	// SlotNumberByRegistrationNumberAction - Action to return slot number of car by registration number
	SlotNumberByRegistrationNumberAction = "slot_number_for_registration_number"

	// ExitAction - Action to exit
	ExitAction = "exit"
)
