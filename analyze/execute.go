package analyze

import (
	"os/exec"
	"os"
	"runtime"
	"fmt"
)
//this function executes the console command, represented as an []string(input).
func Execute(input []string) (string, error) {
	var cmd *exec.Cmd = nil
	switch runtime.GOOS {
	case "linux": cmd = exec.Command(input[0], input [1:]...)
	case "macos": cmd = exec.Command("bash", append([]string{"-c"}, input[:]...)...)
	case "windows": cmd = exec.Command("cmd", append([]string{"/C"}, input[:]...)...)
	default:
		{
			fmt.Println("Unsupportable OS")
			os.Exit(1)
		}
	}
	out, err := cmd.CombinedOutput()
	return string(out), err
}