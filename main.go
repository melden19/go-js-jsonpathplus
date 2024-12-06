package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/melden19/go-js-demo-bridge/execjs"
)

func main() {
	file := execjs.GetJSONFileContent()

	start := time.Now()

	for i := 0; i < 5000; i++ {
		result, err := execjs.ExecJSONPathWithGoJa("$.global.dash-dash.underscore_underscore", string(file))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err)
			os.Exit(1)
		}
		fmt.Printf("[%v] Result of ExecJSONPath: %s\n", i, result)
		PrintMemUsage()
	}

	elapsed := time.Since(start)
	fmt.Printf("Duration: %s\n", elapsed)
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
