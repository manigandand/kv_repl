package store

import (
	"vmware/command"
)

// Write sets/updates the value for the given key
func (s *Store) Write(cmd *command.Command) error {
	s.mem[cmd.Key] = cmd.Value
	return nil
}
