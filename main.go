package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"vmware/command"
	"vmware/store"
	"vmware/utils"
)

var (
	shell *IShell
)

// InitIshell inits the conn
func init() {
	shell = &IShell{
		Prompt: DefaultPrompt,
	}
}

// configure stdin and stdout
func main() {
	store.InitStore()
	fmt.Printf("K/V REPL started...\n\n")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		fmt.Println("Exiting...")
		os.Exit(0)
	}()

	// interactive mode, read from stdin
	// and write to stdout
	reader := bufio.NewReader(os.Stdin)
	for {
		shell.ShowPrompt()

		cmdInputStr, err := reader.ReadString(utils.EndLineDelim)
		if err != nil {
			fmt.Println("error reading input: ", err)
			break
		}
		cmdInputStr = strings.TrimRight(cmdInputStr, utils.NewLineDelim)

		shell.RecordHistory(cmdInputStr)

		cmd, err := command.New(cmdInputStr)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if cmd.IsQuit() {
			fmt.Println("Exiting...")
			os.Exit(0)
			break
		}

		if err := store.ExecCmd(cmd); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
