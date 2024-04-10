package main

import (
	"os"
	"syscall"
	"time"
)

func init() {
	rusage = rusageLinux{}
}

type rusageLinux struct{}

// maxRss returns RSS usage in kB.
// On Linux this is what getrusage(2) returns.
func (rusageLinux) stats(pst *os.ProcessState) (rusageStats, bool) {
	rusage, ok := pst.SysUsage().(*syscall.Rusage)
	if !ok {
		return rusageStats{}, false
	}
	return rusageStats{
		user:   time.Duration(rusage.Utime.Nano()),
		sys:    time.Duration(rusage.Stime.Nano()),
		maxRss: rusage.Maxrss,
		minFlt: rusage.Minflt,
		majFlt: rusage.Majflt,
	}, true
}
