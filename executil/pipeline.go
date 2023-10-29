package executil

import (
	"fmt"
	"io"
	"os/exec"
)

// NewPipeline create a new pipeline, optionally with commands
func NewPipeline(cmds ...*exec.Cmd) *Pipeline {
	return &Pipeline{
		stdout: nil,
		cmds:   cmds,
	}
}

// Pipeline represents a series of commands linked by their stdout and stdin
type Pipeline struct {
	stdout io.ReadCloser
	cmds   []*exec.Cmd
}

// Append adds a new command to the tail end of the pipeline, connecting its
// stdin to the stdout of the previous tail
func (p *Pipeline) Append(cmd *exec.Cmd) (err error) {
	if p.stdout != nil {
		cmd.Stdin = p.stdout
	}

	if p.stdout, err = cmd.StdoutPipe(); err != nil {
		return fmt.Errorf("failed to get stdout for %s: %s", cmd.String(), err)
	}

	p.cmds = append(p.cmds, cmd)

	return nil
}

// Execute starts each command in the pipeline and waits for their completion
// in reverse order, capturing the stdout of the last command
func (p *Pipeline) Execute() (out []byte, err error) {
	// start commands
	for _, cmd := range p.cmds {
		if err = cmd.Start(); err != nil {
			return nil, fmt.Errorf("failed to start cmd %s: %s", cmd.String(), err)
		}
	}

	// start reading output from the tail-end of the pipe
	if out, err = io.ReadAll(p.stdout); err != nil {
		return nil, fmt.Errorf("failed to get pipeline tail stdout: %s", err)
	}

	// wait for commands to finish in reverse
	for i := len(p.cmds) - 1; i > -1; i-- {
		cmd := p.cmds[i]
		if err := cmd.Wait(); err != nil {
			return nil, fmt.Errorf("failed to wait for cmd %s: %s", cmd.String(), err)
		}
	}

	return out, nil
}
