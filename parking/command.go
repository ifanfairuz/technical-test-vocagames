package parking

import "strings"

type Command struct {
	command string
	args    []string
}

// Parse command string to Command struct
func parseCommand(command string) Command {
	args := strings.SplitN(command, " ", 3)

	return Command{
		command: strings.ToLower(args[0]),
		args:    args[1:],
	}
}
