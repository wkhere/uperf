package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"sync/atomic"
	"time"

	"github.com/wkhere/dtf"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		os.Exit(0)
	}

	intr := new(atomic.Bool)
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)
	go func() {
		<-sigc
		intr.Store(true)
	}()

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	t0 := time.Now()
	err := cmd.Run()
	t1 := time.Now()

	if err != nil {
		// errors.As not needed as the ExitError shouldn't be wrapped
		if _, ok := err.(*exec.ExitError); !ok {
			die(1, err)
		}
	}

	printStats(os.Stderr, t1.Sub(t0), cmd.ProcessState, intr.Load())
	os.Exit(cmd.ProcessState.ExitCode())
}

func printStats(w *os.File, wall time.Duration, pst *os.ProcessState, intr bool) {
	title := "uperf: "
	if intr {
		title = "uperf (interrupted): "
	}
	fmt.Fprint(w, title, dtf.Fmt(wall), " total")
	if stats, ok := rusage.stats(pst); ok {
		fmt.Fprint(w, ", ", stats)
	}
	fmt.Fprintln(w)
}

func die(code int, err error) {
	fmt.Fprintln(os.Stderr, "utime:", err)
	os.Exit(code)
}
