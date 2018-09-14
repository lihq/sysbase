package helpers

import (
	"context"
	"os"
	"os/exec"
	"time"
)

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
