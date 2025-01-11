uperf
=====

`uperf` runs given command and prints timings plus memory stats (rss, minor/major faults) portably,
for Linux & Darwin.

```
% uperf ls *.go
main.go  rusage_darwin.go  rusage.go  rusage_linux.go
uperf: 3.252ms total, 0s user, 3.004ms sys, 2948k RSS, 92/0 flt
```

