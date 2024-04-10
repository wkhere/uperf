package main

import (
	"os"
	"syscall"
	"time"
)

func init() {
	rusage = rusageDarwin{}
}

type rusageDarwin struct{}

// maxRss returns RSS usage in kB.
// On Darwin getrusage(2) returns bytes count.
func (rusageDarwin) stats(pst *os.ProcessState) (rusageStats, bool) {
	rusage, ok := pst.SysUsage().(*syscall.Rusage)
	if !ok {
		return rusageStats{}, false
	}
	return rusageStats{
		user:   time.Duration(rusage.Utime.Nano()),
		sys:    time.Duration(rusage.Stime.Nano()),
		maxRss: rusage.Maxrss / 1024,
		minFlt: rusage.Minflt,
		majFlt: rusage.Majflt,
	}, true
}
