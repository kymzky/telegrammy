package utils

import (
	"bytes"
	"os/exec"
)

type Executor struct{}

func NewExecutor() *Executor {
	return &Executor{}
}

func (e *Executor) Execute(command string) (*ExecutionResult, error) {
	args := splitCommand(command)
	if len(args) == 0 {
		return nil, ErrEmptyCommand
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = &bytes.Buffer{}
	cmd.Stderr = &bytes.Buffer{}

	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	stdoutBytes := cmd.Stdout.(*bytes.Buffer).Bytes()
	stderrBytes := cmd.Stderr.(*bytes.Buffer).Bytes()

	return &ExecutionResult{
		Stdout: string(stdoutBytes),
		Stderr: string(stderrBytes),
		Code:   0,
	}, nil
}

func splitCommand(command string) []string {
	args := make([]string, 0)
	currentArg := ""
	inQuotes := false

	for _, char := range command {
		switch char {
		case '"':
			inQuotes = !inQuotes
		case ' ', '\t':
			if inQuotes {
				currentArg += string(char)
			} else if currentArg != "" {
				args = append(args, currentArg)
				currentArg = ""
			}
		default:
			currentArg += string(char)
		}
	}

	if currentArg != "" {
		args = append(args, currentArg)
	}

	return args
}

type ExecutionResult struct {
	Stdout string
	Stderr string
	Code   int
}

var ErrEmptyCommand = &Error{Msg: "empty command provided"}

type Error struct {
	Msg string
}

func (e *Error) Error() string {
	return e.Msg
}
