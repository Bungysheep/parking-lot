package commandmanager

import (
	"testing"
	"utils/constant"
	"utils/message"
)

func TestCreateCommandManager(t *testing.T) {
	cm := NewCommandManager()

	if cm == nil {
		t.Errorf("Expected to have a new Command Manager.")
	}

	if len(cm.Commands) != 7 {
		t.Errorf("Expected to have %d commands, but got %d", 7, len(cm.Commands))
	}

	if cm.Commands[constant.CreateParkingLotAction] == nil {
		t.Errorf("Expected to have %s commands.", constant.CreateParkingLotAction)
	}
	if cm.Commands[constant.ParkAction] == nil {
		t.Errorf("Expected to have %s commands.", constant.ParkAction)
	}
	if cm.Commands[constant.LeaveAction] == nil {
		t.Errorf("Expected to have %s commands.", constant.LeaveAction)
	}
	if cm.Commands[constant.StatusAction] == nil {
		t.Errorf("Expected to have %s commands.", constant.StatusAction)
	}
	if cm.Commands[constant.RegistrationNumberOfCarsByColourAction] == nil {
		t.Errorf("Expected to have %s commands.", constant.RegistrationNumberOfCarsByColourAction)
	}
	if cm.Commands[constant.SlotNumberOfCarsByColourAction] == nil {
		t.Errorf("Expected to have %s commands.", constant.SlotNumberOfCarsByColourAction)
	}
	if cm.Commands[constant.SlotNumberByRegistrationNumberAction] == nil {
		t.Errorf("Expected to have %s commands.", constant.SlotNumberByRegistrationNumberAction)
	}
}

func TestCommandManagerWithoutAction(t *testing.T) {
	cm := NewCommandManager()

	if cm == nil {
		t.Errorf("Expected to have a new Command Manager.")
	}

	_, err := cm.Execute("")

	if err.Error() != message.ActionIsInvalid("") {
		t.Errorf("Expected to have an error saying: %s, but got error: %s", message.ActionIsInvalid(""), err.Error())
	}

	_, err = cm.Execute("ABC")

	if err.Error() != message.CommandIsInvalid("abc") {
		t.Errorf("Expected to have an error saying: %s, but got error: %s", message.CommandIsInvalid("abc"), err.Error())
	}
}

func TestCommandManagerWithInvalidAction(t *testing.T) {
	cm := NewCommandManager()

	if cm == nil {
		t.Errorf("Expected to have a new Command Manager.")
	}

	_, err := cm.Execute("create_parking 5")

	if err.Error() != message.CommandIsInvalid(cm.Action) {
		t.Errorf("Expected to have an error saying: %s, but got error: %s", message.CommandIsInvalid(cm.Action), err.Error())
	}
}

func TestCOmmandManagerWithValidAction(t *testing.T) {
	cm := NewCommandManager()

	if cm == nil {
		t.Errorf("Expected to have a new Command Manager.")
	}

	result, err := cm.Execute("create_parking_lot 5")
	if result != message.ParkingCreated(5) {
		t.Errorf("Expected to have result %s, but got %s", message.ParkingCreated(5), result)
	}

	if err != nil {
		t.Errorf("Expected to have no error, but got error: %s", err.Error())
	}
}

func TestCOmmandManagerWithInvalidActionArgs(t *testing.T) {
	cm := NewCommandManager()

	if cm == nil {
		t.Errorf("Expected to have a new Command Manager.")
	}

	_, err := cm.Execute("create_parking_lot a a")
	if err.Error() != message.ParameterIsInvalid("a a") {
		t.Errorf("Expected to have an error saying: %s, but got error: %s", message.ParameterIsInvalid("a a"), err.Error())
	}

	_, err = cm.Execute("create_parking_lot 0")
	if err.Error() != message.ParkingCapacityLessThanMinimumCapacity() {
		t.Errorf("Expected to have an error saying: %s, but got error: %s", message.ParkingCapacityLessThanMinimumCapacity(), err.Error())
	}
}
