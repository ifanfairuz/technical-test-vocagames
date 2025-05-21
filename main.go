package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ifanfairuz/technical-test-vocagames/parking"
)

func printHelp() {
	fmt.Print("Usage: parkinglot <command_file> \n\n")
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	// Check if the file exists
	commandPath := os.Args[1]
	if _, err := os.Stat(commandPath); os.IsNotExist(err) {
		printHelp()
		log.Fatalln(commandPath, "file does not exist")
	}

	// Open the file
	commandFile, err := os.Open(commandPath)
	if err != nil {
		log.Fatalln("error opening file: ", err)
	}
	defer commandFile.Close() // Close the file when we're done

	parking.RunCommandFile(commandFile, os.Stdout)
}
