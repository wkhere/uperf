package main

import (
	"fmt"
	"os"
	"time"

	"github.com/wkhere/dtf"
)

type rusageStats struct {
	user, sys      time.Duration
	maxRss         int64 // RSS usage in kB
	minFlt, majFlt int64
}

func (m rusageStats) String() string {
	return fmt.Sprintf(
		"%s user, %s sys, %dk RSS, %d/%d flt",
		dtf.Fmt(m.user), dtf.Fmt(m.sys),
		m.maxRss, m.minFlt, m.majFlt,
	)
}

type rusageGetter interface {
	stats(*os.ProcessState) (rusageStats, bool)
}

var rusage rusageGetter = noRusage{}

type noRusage struct{}

func (noRusage) stats(*os.ProcessState) (rusageStats, bool) {
	return rusageStats{}, false
}
