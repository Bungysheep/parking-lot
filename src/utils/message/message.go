package message

import "fmt"

// ParameterIsInvalid message
func ParameterIsInvalid(param string) string {
	return fmt.Sprintf("Parameter '%s' is invalid", param)
}

// ActionIsInvalid message
func ActionIsInvalid(action string) string {
	return fmt.Sprintf("Action '%s' is invalid", action)
}

// CommandIsInvalid message
func CommandIsInvalid(command string) string {
	return fmt.Sprintf("Command '%s' is invalid", command)
}

// ParkingCapacityLessThanMinimumCapacity message
func ParkingCapacityLessThanMinimumCapacity() string {
	return fmt.Sprintf("Parking Capacity is less than Minimum Capacity")
}

// ParkingCreated message
func ParkingCreated(capacity int) string {
	return fmt.Sprintf("Created a parking lot with %d slots", capacity)
}

// ParkingIsFull message
func ParkingIsFull() string {
	return fmt.Sprintf("Sorry, parking lot is full")
}

// SlotIsInvalid message
func SlotIsInvalid(slotNbr int) string {
	return fmt.Sprintf("Slot '%d' is invalid", slotNbr)
}

// NoCarParkedAtSlot message
func NoCarParkedAtSlot(slotNbr int) string {
	return fmt.Sprintf("No Car parked at Slot '%d'", slotNbr)
}

// CarEntered message
func CarEntered(slotNbr int) string {
	return fmt.Sprintf("Allocated slot number: %d", slotNbr)
}

// CarExit message
func CarExit(slotNbr int) string {
	return fmt.Sprintf("Slot number %d is free", slotNbr)
}

// NotFound message
func NotFound() string {
	return fmt.Sprintf("Not found")
}

func NoDataFound() string {
	return fmt.Sprintf("No Data Found")
}
