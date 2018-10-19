package helpers

import (
	"context"
	"log"
	"os"
	"os/exec"
	"time"
)

// ExecCommand runs shell command with timeout and without ENV COLUMNS=80 limit.
func ExecCommand(cmdString string, timeoutInSeconds int) (out []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutInSeconds)*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "/bin/bash", "-c", cmdString)
	cmd.Env = append(os.Environ(),
		"COLUMNS=512", // overwrite montherfuck monkey hard-coded COLUMNS=80 in golang by default
	)
	out, err = cmd.Output()
	return
}

// ExecCommandExitAfterTimeout runs shell command exit main process if timeout.
// NOTICE: context.WithTimeout doesn't works as expected,
// if cmdString contains pipe command `df -h |grep /dev` and get NAS failed may be dead-wait.
func ExecCommandQuitAfterTimeout(cmdString string, timeoutInSeconds int) (out []byte, err error) {
	cmd := exec.Command("/bin/bash", "-c", cmdString)
	cmd.Env = append(os.Environ(),
		"COLUMNS=512",
	)

	timer := time.AfterFunc(3*time.Second, func() {
		cmd.Process.Kill()
		log.Println("quit main process avoid dead-wait.")
		os.Exit(1)
	})
	out, err = cmd.Output()
	timer.Stop()
	return
}
