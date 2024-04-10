package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	t0 := time.Now()
	err := cmd.Run()
	t1 := time.Now()

	if err != nil {
		var e *exec.ExitError
		if !errors.As(err, &e) {
			die(1, err)
		}
	}

	fmt.Fprintln(os.Stderr, "utime:", t1.Sub(t0))
	os.Exit(cmd.ProcessState.ExitCode())
}

func die(code int, err error) {
	fmt.Fprintln(os.Stderr, "utime:", err)
	os.Exit(1)
}
