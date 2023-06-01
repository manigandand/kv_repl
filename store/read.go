package store

import (
	"fmt"
	"kv_repl/command"
)

// Read reads the value for the given key
func (s *Store) Read(cmd *command.Command, ingoreTrans bool) error {
	if s.hasActiveTransaction() && ingoreTrans {
		// read from the current transaction mem
		return s.transactions[len(s.transactions)-1].Read(cmd)
	}

	val, ok := s.mem[cmd.Key]
	if !ok {
		return fmt.Errorf("key not found: %s", cmd.Key)
	}

	fmt.Println(val)
	return nil
}
