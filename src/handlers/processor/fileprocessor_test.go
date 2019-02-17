package processor

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"utils/constant"
)

var inputs = []byte(
	`create_parking_lot 6
	park KA-01-HH-1234 White
	park KA-01-HH-9999 White
	park KA-01-BB-0001 Black
	park KA-01-HH-7777 Red
	park KA-01-HH-2701 Blue
	park KA-01-HH-3141 Black
	leave 4
	status
	park KA-01-P-333 White
	park DL-12-AA-9999 White
	registration_numbers_for_cars_with_colour White
	slot_numbers_for_cars_with_colour White
	slot_number_for_registration_number KA-01-HH-3141
	slot_number_for_registration_number MH-04-AY-1111`)

func TestRunFileProcessor(t *testing.T) {
	if err := ioutil.WriteFile("file_inputs_test.txt", inputs, 0666); err != nil {
		t.Fatal(err)
	}
	defer os.Remove("file_inputs_test.txt")

	fp := NewFileProcessor("file_inputs_test.txt")

	if fp == nil {
		t.Error("Expected to have a new File Processor.")
	}

	resp := []string{
		"Created a parking lot with 6 slots",
		"Allocated slot number: 1",
		"Allocated slot number: 2",
		"Allocated slot number: 3",
		"Allocated slot number: 4",
		"Allocated slot number: 5",
		"Allocated slot number: 6",
		"Slot number 4 is free",
		fmt.Sprintf("%-10s%-18s%-10s", "Slot No.", "Registration No", "Colour"),
		fmt.Sprintf("%-10d%-18s%-10s", 1, "KA-01-HH-1234", "White"),
		fmt.Sprintf("%-10d%-18s%-10s", 2, "KA-01-HH-9999", "White"),
		fmt.Sprintf("%-10d%-18s%-10s", 3, "KA-01-BB-0001", "Black"),
		fmt.Sprintf("%-10d%-18s%-10s", 5, "KA-01-HH-2701", "Blue"),
		fmt.Sprintf("%-10d%-18s%-10s", 6, "KA-01-HH-3141", "Black"),
		"Allocated slot number: 4",
		"Sorry, parking lot is full",
		"KA-01-HH-1234, KA-01-HH-9999, KA-01-P-333",
		"1, 2, 4",
		"6",
		"Not found",
	}
	expectedResp := strings.Join(resp, constant.NewLine)

	fp.Run()
	if fp.Resp != expectedResp {
		t.Errorf("Expected to have response:\n %s, \nbut got:\n %s\n", expectedResp, fp.Resp)
	}
}
