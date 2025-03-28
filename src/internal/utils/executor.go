package utils

import (
	"bytes"
	"log/slog"
	"os/exec"
)

type Executor struct{}

func NewExecutor() *Executor {
	return &Executor{}
}

func (e *Executor) Execute(shellPath string, command string) string {
	slog.Info("Execute command", "command", command)

	cmd := exec.Command(shellPath, "-c", command)
	cmd.Stdout = &bytes.Buffer{}
	cmd.Stderr = &bytes.Buffer{}

	err := cmd.Start()
	if err != nil {
		slog.Error("Error on command start", "error", err)
		return err.Error()
	}

	err = cmd.Wait()
	if err != nil {
		slog.Error("Error on command wait", "error", err)
		stderr := string(cmd.Stderr.(*bytes.Buffer).Bytes())
		slog.Warn("Command stderr:", "stderr", stderr)
		return stderr
	}

	stdout := string(cmd.Stdout.(*bytes.Buffer).Bytes())
	slog.Info("Command output:", "stdout", stdout)
	return stdout
}

var ErrEmptyCommand = &Error{Msg: "empty command provided"}

type Error struct {
	Msg string
}

func (e *Error) Error() string {
	return e.Msg
}
