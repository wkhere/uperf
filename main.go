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

	printStats(os.Stderr, t1.Sub(t0), cmd.ProcessState)
	os.Exit(cmd.ProcessState.ExitCode())
}

func printStats(w *os.File, wall time.Duration, pst *os.ProcessState) {
	fmt.Fprint(w, "uperf: ", wall, " total")
	if mem, ok := rusage.stats(pst); ok {
		fmt.Fprint(w, ", ", mem)
	}
	fmt.Fprintln(w)
}

func die(code int, err error) {
	fmt.Fprintln(os.Stderr, "utime:", err)
	os.Exit(1)
}
