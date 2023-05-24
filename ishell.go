package main

import (
	"fmt"
	"time"
)

const (
	// DefaultPrompt is string prefix.
	DefaultPrompt = "$ "
)

// IShell is an interactive cli shell.
type IShell struct {
	Prompt  string
	History []*IShellHistory
}

// IShellHistory holds all the commands history
type IShellHistory struct {
	Command   string
	CreatedAt time.Time
}

// ShowPrompt displays the DefaultPrompt, and wait for the next commands
func (s *IShell) ShowPrompt() {
	fmt.Print(s.Prompt)
}

// RecordHistory holds the History for shell session
func (s *IShell) RecordHistory(cmd string) {
	record := &IShellHistory{
		Command:   cmd,
		CreatedAt: time.Now(),
	}
	s.History = append(s.History, record)
}
