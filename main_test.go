package main_test

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"testing"

	"github.com/ifanfairuz/technical-test-vocagames/parking"
)

func TestCommandCase1(t *testing.T) {
	// create virtual stdout
	out := &bytes.Buffer{}

	// Open the file
	commandFile, err := os.Open("command_case1.txt")
	if err != nil {
		log.Fatalln("error opening file: ", err)
	}
	defer commandFile.Close() // Close the file when we're done

	// Run the command file
	parking.RunCommandFile(commandFile, out)

	// Check the output
	expected := "Allocated slot number: 1\n"
	expected += "Allocated slot number: 2\n"
	expected += "Allocated slot number: 3\n"
	expected += "Allocated slot number: 4\n"
	expected += "Allocated slot number: 5\n"
	expected += "Allocated slot number: 6\n"
	expected += "Registration number KA-01-HH-3141 with Slot Number 6 is free with Charge $30\n"
	expected += "Slot No. Registration No.\n"
	expected += "1 KA-01-HH-1234\n"
	expected += "2 KA-01-HH-9999\n"
	expected += "3 KA-01-BB-0001\n"
	expected += "4 KA-01-HH-7777\n"
	expected += "5 KA-01-HH-2701\n"
	expected += "Allocated slot number: 6\n"
	expected += "Sorry, parking lot is full\n"
	expected += "Registration number KA-01-HH-1234 with Slot Number 1 is free with Charge $30\n"
	expected += "Registration number KA-01-BB-0001 with Slot Number 3 is free with Charge $50\n"
	expected += "Registration number DL-12-AA-9999 not found\n"
	expected += "Allocated slot number: 1\n"
	expected += "Allocated slot number: 3\n"
	expected += "Sorry, parking lot is full\n"
	expected += "Slot No. Registration No.\n"
	expected += "1 KA-09-HH-0987\n"
	expected += "2 KA-01-HH-9999\n"
	expected += "3 CA-09-IO-1111\n"
	expected += "4 KA-01-HH-7777\n"
	expected += "5 KA-01-HH-2701\n"
	expected += "6 KA-01-P-333\n"

	result := out.String()
	if result != expected {
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
				strDebug += fmt.Sprintf("false => nil -> %q\n", resultLines[i])
			} else {
				strDebug += fmt.Sprintf("false => %q -> nil\n", expectedLines[i])
			}
		}

		t.Error(strDebug)
	}
}

func BenchmarkParking(b *testing.B) {
	parkingLot := parking.NewParkingLot(b.N)

	for i := 1; i <= b.N; i++ {
		if i%9 == 0 {
			parkingLot.Status()
		} else if i%7 == 0 {
			past := rand.Intn(i-2) + 1
			parkingLot.Leave(fmt.Sprintf("KA-01-HH-%04d", past), past)
		} else {
			parkingLot.Park(fmt.Sprintf("KA-01-HH-%04d", i))
		}
	}
}
