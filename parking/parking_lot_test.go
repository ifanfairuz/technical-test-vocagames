package parking_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ifanfairuz/technical-test-vocagames/parking"
)

func debugResult(expected, result string) string {
	strDebug := "\nEqual => Expected -> got\n"

	expectedLines := strings.Split(expected, "\n")
	resultLines := strings.Split(result, "\n")
	for i := 0; i < max(len(expectedLines), len(resultLines)); i++ {
		if i < len(expectedLines) && i < len(resultLines) {
			equal := "ERR"
			if expectedLines[i] == resultLines[i] {
				equal = "OK"
			}
			strDebug += fmt.Sprintf("%s => %q -> %q\n", equal, expectedLines[i], resultLines[i])
		} else if i < len(resultLines) {
			strDebug += fmt.Sprintf("ERR => nil -> %q\n", resultLines[i])
		} else {
			strDebug += fmt.Sprintf("ERR => %q -> nil\n", expectedLines[i])
		}
	}

	return strDebug
}

// Test create parking
func TestCreateParkingLot(t *testing.T) {
	parkingLot := parking.NewParkingLot(6)
	size := parkingLot.Size()
	if size != 6 {
		t.Errorf("Expected parking lot capacity to be 6, got %d", size)
	}
}

// Test park
func TestPark(t *testing.T) {
	parkingLot := parking.NewParkingLot(6)

	result := parkingLot.Park("KA-01-HH-1234")

	// check output
	expected := "Allocated slot number: 1\n"
	if result != expected {
		t.Error(debugResult(expected, result))
	}
}

// Test park no space
func TestParkNoSpace(t *testing.T) {
	parkingLot := parking.NewParkingLot(2)

	result := parkingLot.Park("KA-01-HH-1234")
	result += parkingLot.Park("KA-02-HH-1234")
	result += parkingLot.Park("KA-03-HH-1235")

	// check output
	expected := "Allocated slot number: 1\n"
	expected += "Allocated slot number: 2\n"
	expected += "Sorry, parking lot is full\n"
	if result != expected {
		t.Error(debugResult(expected, result))
	}
}

// Test leave
func TestLeave(t *testing.T) {
	parkingLot := parking.NewParkingLot(6)

	result := parkingLot.Park("KA-01-HH-1234")
	result += parkingLot.Leave("KA-01-HH-1234", 4)

	// check output
	expected := "Allocated slot number: 1\n"
	expected += "Registration number KA-01-HH-1234 with Slot Number 1 is free with Charge $30\n"
	if result != expected {
		t.Error(debugResult(expected, result))
	}
}

// Test unregistered leave
func TestUnregisteredLeave(t *testing.T) {
	parkingLot := parking.NewParkingLot(6)

	result := parkingLot.Park("KA-01-HH-1234")
	result += parkingLot.Leave("KA-01-HH-1235", 4)

	// check output
	expected := "Allocated slot number: 1\n"
	expected += "Registration number KA-01-HH-1235 not found\n"
	if result != expected {
		t.Error(debugResult(expected, result))
	}
}

// Test park leave and park again
func TestParkLeavePark(t *testing.T) {
	parkingLot := parking.NewParkingLot(4)

	result := parkingLot.Park("KA-01-HH-1234")
	result += parkingLot.Park("KA-02-HH-1234")
	result += parkingLot.Park("KA-03-HH-1234")
	result += parkingLot.Park("KA-04-HH-1234")
	result += parkingLot.Leave("KA-02-HH-1234", 2)
	result += parkingLot.Leave("KA-04-HH-1234", 3)
	result += parkingLot.Park("KA-05-HH-1234")
	result += parkingLot.Leave("KA-03-HH-1234", 4)
	result += parkingLot.Park("KA-06-HH-1234")
	result += parkingLot.Park("KA-07-HH-1234")
	result += parkingLot.Park("KA-08-HH-1234")

	// check output
	expected := "Allocated slot number: 1\n"
	expected += "Allocated slot number: 2\n"
	expected += "Allocated slot number: 3\n"
	expected += "Allocated slot number: 4\n"
	expected += "Registration number KA-02-HH-1234 with Slot Number 2 is free with Charge $10\n"
	expected += "Registration number KA-04-HH-1234 with Slot Number 4 is free with Charge $20\n"
	expected += "Allocated slot number: 2\n"
	expected += "Registration number KA-03-HH-1234 with Slot Number 3 is free with Charge $30\n"
	expected += "Allocated slot number: 3\n"
	expected += "Allocated slot number: 4\n"
	expected += "Sorry, parking lot is full\n"
	if result != expected {
		t.Error(debugResult(expected, result))
	}
}

// Test status
func TestStatus(t *testing.T) {
	parkingLot := parking.NewParkingLot(6)

	result := parkingLot.Park("KA-01-HH-1234")
	result += parkingLot.Park("KA-02-HH-1234")
	result += parkingLot.Park("KA-03-HH-1234")
	result += parkingLot.Leave("KA-02-HH-1234", 2)
	result += parkingLot.Status()

	// check output
	expected := "Allocated slot number: 1\n"
	expected += "Allocated slot number: 2\n"
	expected += "Allocated slot number: 3\n"
	expected += "Registration number KA-02-HH-1234 with Slot Number 2 is free with Charge $10\n"
	expected += "Slot No. Registration No.\n"
	expected += "1 KA-01-HH-1234\n"
	expected += "3 KA-03-HH-1234\n"
	if result != expected {
		t.Error(debugResult(expected, result))
	}
}
