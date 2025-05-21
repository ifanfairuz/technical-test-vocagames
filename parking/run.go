package parking

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

// RunCommandFile runs a set of commands from a file
func RunCommandFile(file *os.File, out io.Writer) {
	// Read commands from the file
	scanner := bufio.NewScanner(file)
	currentLine := 1

	// check begin command
	if !scanner.Scan() {
		log.Fatalln("Error on command file line", currentLine, ": Empty file")
	}

	// check begin command, must be create_parking_lot <capacity>
	beginCommand := parseCommand(scanner.Text())
	if beginCommand.command != "create_parking_lot" || len(beginCommand.args) < 1 {
		log.Fatalln("Error on command file line", currentLine, ": The first command must be create_parking_lot <capacity>")
	}

	// parse capacity
	capacity, err := strconv.Atoi(beginCommand.args[0])
	if err != nil {
		log.Fatalln("Error on command file line", currentLine, ": capacity must be a integer")
	}

	// create parking lot and run commands
	parkingLot := NewParkingLot(capacity)
	for scanner.Scan() {
		currentLine++
		res, err := parkingLot.RunCommand(scanner.Text())
		if err != nil {
			log.Fatalln("Error on command file line", currentLine, ":", err)
		}

		out.Write([]byte(res))
	}
}
