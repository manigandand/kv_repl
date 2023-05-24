package store

import (
	"fmt"
	"vmware/command"
)

// Read reads the value for the given key
func (s *Store) Read(cmd *command.Command) error {
	val, ok := s.mem[cmd.Key]
	if !ok {
		return fmt.Errorf("key not found: %s", cmd.Key)
	}

	fmt.Println(val)
	return nil
}
