package execute

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
)

// Run function runs binary with Context and arguments. Result: (string stdout, string stderr)
func Run(context Context) (string, string) {
	args := []string{}
	for _, element := range context.Args {
		if len(element.Key) > 0 {
			args = append(args, fmt.Sprintf("%s=%s", element.Key, element.Value)) // TODO: parse delimeters for each engine
		} else {
			args = append(args, element.Value)
		}
	}
	cmd := exec.Command(context.Binary, args...)
	if len(context.WorkingDirectory) > 0 {
		cmd.Dir = context.WorkingDirectory
	} else {
		dir, _ := os.Getwd()
		cmd.Dir = dir
	}
	if len(context.Stdin) > 0 {
		stdin, _ := cmd.StdinPipe()
		defer stdin.Close()
		io.WriteString(stdin, context.Stdin)
	}
	var stderr, stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Start(); err != nil {
		return "", context.Binary + " crashed with: " + err.Error()
	}

	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			status, _ := exiterr.Sys().(syscall.WaitStatus)
			if status.ExitStatus() != context.SuccessCode {
				return "", context.Binary + " crashed with: " + err.Error()
			}
		}
	}

	return string(stdout.Bytes()), string(stderr.Bytes())
}
