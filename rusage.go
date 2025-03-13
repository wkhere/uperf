package main

import (
	"os"
	"time"
)

type rusageStats struct {
	user, sys      time.Duration
	maxRss         int64 // RSS usage in kB
	minFlt, majFlt int64
}

type rusageGetter interface {
	stats(*os.ProcessState) (rusageStats, bool)
}

var rusage rusageGetter = noRusage{}

type noRusage struct{}

func (noRusage) stats(*os.ProcessState) (rusageStats, bool) {
	return rusageStats{}, false
}
