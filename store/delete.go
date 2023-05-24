package store

import (
	"fmt"
	"vmware/command"
)

// Delete deletes the value for the given key
func (s *Store) Delete(cmd *command.Command) error {
	if _, ok := s.mem[cmd.Key]; !ok {
		return fmt.Errorf("key not found: %s", cmd.Key)
	}
	delete(s.mem, cmd.Key)

	return nil
}
