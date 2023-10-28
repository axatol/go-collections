package executil

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func Pipeline(cmds ...*exec.Cmd) (output []byte, err error) {
	var stdout io.ReadCloser

	// build up stdin->stdout connections
	for _, cmd := range cmds {
		if stdout != nil {
			cmd.Stdin = stdout
		}

		if stdout, err = cmd.StdoutPipe(); err != nil {
			return nil, fmt.Errorf("failed to get stdout for %s: %s", cmd.String(), err)
		}
	}

	// start commands
	for _, cmd := range cmds {
		if err = cmd.Start(); err != nil {
			return nil, fmt.Errorf("failed to start cmd %s: %s", cmd.String(), err)
		}
	}

	// read final output
	if output, err = io.ReadAll(stdout); err != nil {
		err = errors.Join(err, fmt.Errorf("failed to get pipeline tail stdout: %s", err))
	}

	// wait for commands to finish in reverse
	for i := len(cmds) - 1; i > -1; i-- {
		cmd := cmds[i]
		if err := cmd.Wait(); err != nil {
			return nil, fmt.Errorf("failed to wait for cmd %s: %s", cmd.String(), err)
		}
	}

	return output, err
}
