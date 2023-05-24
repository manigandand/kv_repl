package command

import (
	"fmt"
	"strings"
	"vmware/utils"
)

// CMD is a custom type for all available commands
type CMD string

// all available commands constants
const (
	CMDRead   CMD = "read"
	CMDWrite  CMD = "write"
	CMDDelete CMD = "delete"
	CMDStart  CMD = "start"
	CMDCommit CMD = "commit"
	CMDAbort  CMD = "abort"
	CMDQuit   CMD = "quit"
)

// New process the command input string and validates the command & arguments list
// are valid or not and returns the command object
func New(cmdInputStr string) (*Command, error) {
	cmd := &Command{}
	cmdInputStr = strings.TrimSpace(cmdInputStr)
	if cmdInputStr == "" {
		return nil, fmt.Errorf("command shouldn't be empty")
	}

	cmdInputList, err := utils.SplitCmdArguments(cmdInputStr)
	if err != nil {
		return nil, err
	}
	if len(cmdInputList) < 1 {
		return nil, fmt.Errorf("invalid command: %s", cmdInputStr)
	}

	cmd.Name = CMD(strings.ToLower(cmdInputList[0]))

	// set key if exists
	if len(cmdInputList) > 1 {
		cmd.Key = cmdInputList[1]
		cmd.Arguments = append(cmd.Arguments, cmdInputList[1])
	}

	// set value if exists
	if len(cmdInputList) > 2 {
		cmd.Value = cmdInputList[2]
		cmd.Arguments = append(cmd.Arguments, cmdInputList[2])
	}

	// validate the inputs
	if err := cmd.Ok(); err != nil {
		return nil, err
	}

	// set cmd store connection
	// setStoreConnection(cmd)

	return cmd, nil
}

// String returns the string value of CMD
func (c CMD) String() string {
	return string(c)
}

// validCommands is a map of all valid commands
var validCommands = map[CMD]bool{
	CMDRead:   true,
	CMDWrite:  true,
	CMDDelete: true,
	CMDStart:  true,
	CMDCommit: true,
	CMDAbort:  true,
	CMDQuit:   true,
}

// CMDArgumentLength holds the exact arguments length to read for commands
var CMDArgumentLength = map[CMD]int{
	CMDRead:   1, // key
	CMDWrite:  2, // key and value
	CMDDelete: 1, // key
	CMDStart:  0,
	CMDCommit: 0,
	CMDAbort:  0,
	CMDQuit:   0,
}

// Command holds the command properties
type Command struct {
	Name      CMD
	Arguments []string
	Key       string
	Value     string
	// Connection   CMDStore
}

// Ok validate the command object by name and arguments list
func (cmd *Command) Ok() error {
	if _, ok := validCommands[cmd.Name]; !ok {
		return fmt.Errorf("%s is invalid command", cmd.Name)
	}

	if CMDArgumentLength[cmd.Name] == 0 && len(cmd.Arguments) > 0 {
		return fmt.Errorf("%s doesn't require arguments", cmd.Name)
	}

	if len(cmd.Arguments) > CMDArgumentLength[cmd.Name] {
		return fmt.Errorf("%s requires only %d arguments", cmd.Name, CMDArgumentLength[cmd.Name])
	}

	if len(cmd.Arguments) != CMDArgumentLength[cmd.Name] {
		return fmt.Errorf("%s requires %d arguments", cmd.Name, CMDArgumentLength[cmd.Name])
	}

	return nil
}

// IsQuit returns true if the command is quit
func (cmd *Command) IsQuit() bool {
	return cmd.Name == CMDQuit
}
