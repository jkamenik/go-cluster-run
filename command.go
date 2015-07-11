package cluster

import (
	"io"
	"os"
)

type Cmd struct {
	Path string
	Args []string

	Nodes   []string
	SshArgs []string

	Stderr io.Writer
	Stdout io.Writer

	// errors in the form of node
	Errors map[string]error
}

func Command(nodes []string, name string, args ...string) *Cmd {
	cmd := &Cmd{Path: name, Args: args, Nodes: nodes}

	// by default include the options to ignore key and host file
	cmd.SshArgs = []string{""}

	// by default attache to stand outs
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd
}

// Run the command on all nodes serially.  At the end it returns an
// error if any command errors.  All errors can be gotten from the
// cmd.Errors
func RunSerial() error {
	// do something good
}

// Run all the commands in parallel, but wait until they are all
// finished before returning the error status.
func RunParallel() error {
	// do something good
}
