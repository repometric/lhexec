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
			var delimeters = "="
			if len(context.Delimeters) > 0 {
				delimeters = context.Delimeters
			}
			args = append(args, fmt.Sprintf("%s%s%s", element.Key, delimeters, element.Value)) // TODO: parse delimeters for each engine
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

	stdin, stdinErr := cmd.StdinPipe()
	if stdinErr != nil {
		return errorString(context.Binary, stdinErr.Error())
	}

	if len(context.Stdin) > 0 {
		go func() {
			defer stdin.Close()
			io.WriteString(stdin, context.Stdin)
		}()
	}

	var stderr, stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Start(); err != nil {
		return errorString(context.Binary, err.Error())
	}

	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			status, _ := exiterr.Sys().(syscall.WaitStatus)
			if status.ExitStatus() != context.SuccessCode {
				return errorString(context.Binary, err.Error())
			}
		}
	}

	return string(stdout.Bytes()), string(stderr.Bytes())
}

func errorString(command string, err string) (string, string) {
	return "", command + " crashed with: " + err
}
