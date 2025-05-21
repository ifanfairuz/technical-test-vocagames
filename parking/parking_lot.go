package parking

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
)

type ParkingLot struct {
	size     int
	cars     map[int]string
	position map[string]int
	next     []int
}

// NewParkingLot creates a new parking lot
func NewParkingLot(capacity int) ParkingLot {
	return ParkingLot{
		size:     capacity,
		cars:     map[int]string{},
		position: map[string]int{},
		next:     []int{0},
	}
}

// getNearestEmptySlot returns the nearest empty slot
// find the nearest empty slot, return (int)nearest with (bool)available
func (p *ParkingLot) getNearestEmptySlot() (int, bool) {
	pos := p.next[0]

	if pos >= p.size {
		// no empty slot
		return 0, false
	}

	if len(p.next) == 1 {
		p.next[0]++
	} else {
		p.next = p.next[1:]
	}

	return pos, true
}

// Lot returns the capacity of the parking lot
func (p *ParkingLot) Size() int {
	return p.size
}

// Park parks a car in the parking lot
//
//   - If no space available, return "Sorry, parking lot is full"
//   - If space available, return "Allocated slot number: <slot>"
//
// every return value return with a new line at the end of string
func (p *ParkingLot) Park(car string) string {
	pos, ok := p.getNearestEmptySlot()
	if !ok {
		return "Sorry, parking lot is full\n"
	}

	p.cars[pos] = car
	p.position[car] = pos
	return fmt.Sprintf("Allocated slot number: %d\n", pos+1)
}

// Leave a car from the parking lot
//
//   - If the car is not in the parking lot, return "Sorry, car not found"
//   - If the car is in the parking lot, return "Registration number <car> with Slot Number <slot> is free with Charge $<price>"
//
// every return value return with a new line at the end of string
func (p *ParkingLot) Leave(car string, hours int) string {
	if pos, ok := p.position[car]; ok {
		delete(p.cars, pos)
		delete(p.position, car)

		price := 10
		if hours > 2 {
			price += (hours - 2) * 10
		}

		last := len(p.next) - 1
		if p.next[last] == pos+1 {
			p.next[last] = pos
		} else {
			p.next = append(p.next, pos)
			slices.Sort(p.next)
		}

		if p.next[len(p.next)-1] == pos+1 {
			p.next[len(p.next)-1] = pos
		}

		return fmt.Sprintf("Registration number %s with Slot Number %d is free with Charge $%d\n", car, pos+1, price)
	}

	return fmt.Sprintf("Registration number %s not found\n", car)
}

// Status prints the status of the parking lot,
// skip the empty slot
//
// return with a new line at the end of string
//
// example output, with empty in slot 3:
//
//	`Slot No. Registration No.
//	1 KA-01-HH-1234
//	2 KA-01-HH-9999
//	4 KA-01-HH-7777
//	5 KA-01-HH-2701
//	6 KA-01-P-333
//	`
func (p *ParkingLot) Status() string {
	result := "Slot No. Registration No.\n"

	listed := 0
	for pos := 0; pos < p.size; pos++ {
		if car, ok := p.cars[pos]; ok {
			result += fmt.Sprintf("%d %s\n", pos+1, car)

			listed++
			if listed == len(p.cars) {
				break
			}
		}

	}

	return result
}

// RunCommand run a command
//
// run a string command and return the result
// if the command is invalid, return an error
//
// list of command:
//   - park <car_number>
//   - leave <car_number> <hours>
//   - status
//
// error:
//   - invalid command park: car number is required, usage: park <car_number>
//   - invalid command leave: car number and hours are required, usage: leave <car_number> <hours>
//   - invalid command leave: hours must be a integer
//   - invalid command: <command>
func (p *ParkingLot) RunCommand(command string) (string, error) {
	c := parseCommand(command)

	switch c.command {

	// run parking lot commands
	case "park":
		if len(c.args) < 1 {
			return "", errors.New("invalid command park: car number is required, usage: park <car_number>")
		}

		return p.Park(c.args[0]), nil

	// run parking lot leave commands
	case "leave":
		if len(c.args) < 2 {
			return "", errors.New("invalid command leave: car number and hours are required, usage: leave <car_number> <hours>")
		}
		hours, err := strconv.Atoi(c.args[1])
		if err != nil {
			return "", errors.New("invalid command leave: hours must be a integer")
		}

		return p.Leave(c.args[0], hours), nil

	// run parking lot status command
	case "status":
		return p.Status(), nil
	}

	return "", errors.New("invalid command: " + c.command)
}
