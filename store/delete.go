package store

import (
	"fmt"
	"kv_repl/command"
)

// Delete deletes the value for the given key
func (s *Store) Delete(cmd *command.Command, commit bool) error {
	if s.hasActiveTransaction() && !commit {
		// read from the current transaction mem
		return s.transactions[len(s.transactions)-1].Delete(cmd)
	}

	if _, ok := s.mem[cmd.Key]; !ok {
		return fmt.Errorf("key not found: %s", cmd.Key)
	}
	delete(s.mem, cmd.Key)

	return nil
}
