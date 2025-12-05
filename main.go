package main

import (
	"bufio"
	//"bytes"
	"fmt"
	"flag"
	"io"
	"os"
	//"math"
	"runtime"
	"runtime/pprof"
	"time"
)

type versionFunc func(string, io.Writer) error
var versionFuncs = []versionFunc{v1, v2, v3, v4}
var maxGoroutines int

func main() {
	var (
		cpuProfile = flag.String("cpuprofile", "", "write CPU profile to file")
		version = flag.Int("version", len(versionFuncs), "version of solution to run")
		goroutines = flag.Int("goroutines", 0, "no. of goroutines to use for parallelised versions")
		//benchAll = flag.Bool("benchall", false, "benchmark all versions")
	)
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: 1brc-go [-cpuprofile=PROFILE] [-version=N]")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *version < 1 || *version > len(versionFuncs) {
		fmt.Fprintf(os.Stderr, "Invalid version %d\n", *version)
		os.Exit(1)
	}

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(2)
	}
	inputPath := args[0]

	maxGoroutines = *goroutines
	if maxGoroutines == 0 {
		maxGoroutines = runtime.NumCPU()
	}

	st, err := os.Stat(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	size := st.Size()

	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	t0 := time.Now()
	output := bufio.NewWriter(os.Stdout)

	vf := versionFuncs[*version-1]
	err = vf(inputPath, output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	output.Flush()
	elapsed := time.Since(t0)
	fmt.Fprintf(os.Stderr, "Processed %.1fMB in %s\n", float64(size)/(1024*1024), elapsed)
}