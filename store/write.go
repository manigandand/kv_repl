package store

import (
	"kv_repl/command"
)

// Write sets/updates the value for the given key
func (s *Store) Write(cmd *command.Command, commit bool) error {
	if s.hasActiveTransaction() && !commit {
		// read from the current transaction mem
		return s.transactions[len(s.transactions)-1].Write(cmd)
	}

	s.mem[cmd.Key] = cmd.Value
	return nil
}
