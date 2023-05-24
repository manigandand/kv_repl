package store

import (
	"fmt"
	"vmware/command"
)

var (
	// Store holds the store conn interface
	store *Store
)

// InitStore inits the store
func InitStore() {
	store = &Store{
		mem: make(map[string]string),
	}
	fmt.Println("store initialized")
}

// Store holds the in-memory key value store and the transactions state and history
type Store struct {
	// no lock required as the read and write are sequential
	mem map[string]string
}

// ExecCmd executes the command
func ExecCmd(cmd *command.Command) error {
	switch cmd.Name {
	case command.CMDRead:
		return store.Read(cmd)
	case command.CMDWrite:
		return store.Write(cmd)
	case command.CMDDelete:
		return store.Delete(cmd)
	case command.CMDStart:
		// return store.Start(cmd)
	case command.CMDCommit:
		// return store.Commit(cmd)
	case command.CMDAbort:
		// return store.Abort(cmd)
	case command.CMDQuit:
	}
	return nil
}
