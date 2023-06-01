package store

import (
	"fmt"
	"kv_repl/command"
)

func (s *Store) hasActiveTransaction() bool {
	return len(s.transactions) > 0
}

type transaction struct {
	level             int
	mem               map[string]string
	delmem            map[string]string
	actionableCmds    []*command.Command
	parentTransaction *transaction
	root              *Store // root available for first child
}

func (t *transaction) haveReachedFirst() bool {
	return t.level == 1
}

// Read reads the value for the given key
func (t *transaction) Read(cmd *command.Command) error {
	if _, ok := t.delmem[cmd.Key]; ok {
		return fmt.Errorf("key not found: %s", cmd.Key)
	}

	val, ok := t.mem[cmd.Key]
	if !ok {
		// check if the parent transaction has the key available
		// untill we reach the root
		// we need to know the [dl] ---->
		if t.haveReachedFirst() {
			return t.root.Read(cmd, true)
		}

		return t.parentTransaction.Read(cmd)
		// return fmt.Errorf("key not found: %s", cmd.Key)
	}
	fmt.Println(val)

	return nil
}

// Write sets/updates the value for the given key
func (t *transaction) Write(cmd *command.Command) error {
	if _, ok := t.delmem[cmd.Key]; ok {
		// remove the key from delete map
		delete(t.delmem, cmd.Key)
	}

	t.mem[cmd.Key] = cmd.Value

	// add the command to the actionable commands
	t.actionableCmds = append(t.actionableCmds, cmd)

	return nil
}

// Delete deletes the value for the given key
func (t *transaction) Delete(cmd *command.Command) error {
	// add the command to the actionable commands
	// write key1 "val"
	// start
	// read key1
	// del key1 <- it should actally delete
	// read key1 <- read 404
	// abort -> TODO: make a reverse change.... pre val we need to know...
	// ---------------> while del, get the previous key val and store for reverse action
	// ---------------> write, get the previous key val and store for reverse action
	// read key1
	// ------------

	t.actionableCmds = append(t.actionableCmds, cmd)

	// add it to the temp
	t.delmem[cmd.Key] = "" //

	// if _, ok := t.mem[cmd.Key]; !ok {
	// 	// if t.haveReachedFirst() {
	// 	// 	return t.root.Delete(cmd, true)
	// 	// }

	// 	// return t.parentTransaction.Delete(cmd)
	// 	// return fmt.Errorf("key not found: %s", cmd.Key)
	// }
	// delete(t.mem, cmd.Key)

	return nil
}

// Start starts a new transaction
func (s *Store) Start(cmd *command.Command) error {
	trans := &transaction{
		level:          len(s.transactions) + 1,
		mem:            make(map[string]string),
		delmem:         make(map[string]string),
		actionableCmds: make([]*command.Command, 0),
	}

	// assign the parent transaction or the root
	trans.root = s
	if len(s.transactions) != 0 {
		trans.parentTransaction = s.transactions[len(s.transactions)-1]
	}

	s.transactions = append(s.transactions, trans)
	return nil
}

// Commit commits all the actionable commands in the current transaction
// to the parent transaction or to the root store
func (s *Store) Commit(cmd *command.Command) error {
	// check commit without any active transaction
	if !s.hasActiveTransaction() {
		return fmt.Errorf("no active transaction")
	}

	currentTrans := s.transactions[len(s.transactions)-1]

	// commit to the root store
	toCommitRootStore := (len(s.transactions) == 1)

	var parentTrans *transaction
	if !toCommitRootStore {
		parentTrans = s.transactions[len(s.transactions)-2]
	}

	for _, acmd := range currentTrans.actionableCmds {
		if toCommitRootStore {
			// commit to the root store
			// expecting no error here
			switch acmd.Name {
			case command.CMDWrite:
				log(s.Write(acmd, true))
			case command.CMDDelete:
				log(s.Delete(acmd, true))
			}

			continue
		}

		switch acmd.Name {
		case command.CMDWrite:
			log(parentTrans.Write(acmd))
		case command.CMDDelete:
			log(parentTrans.Delete(acmd))
		}
	}

	// remove the current transaction from the stack
	s.transactions = s.transactions[:len(s.transactions)-1]

	return nil
}

func log(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// Abort aborts the transaction, it just removes the current transaction from the stack
func (s *Store) Abort(cmd *command.Command) error {
	// check abort without any active transaction
	if !s.hasActiveTransaction() {
		return fmt.Errorf("no active transaction")
	}

	// reverse change...
	// we store more info of the cmd's previous state

	// remove the current transaction from the stack
	s.transactions = s.transactions[:len(s.transactions)-1]

	return nil
}
