package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/jpsas31/SWE/indexer/parser"
)

// Define flags with default values and descriptions
var filePath = flag.String("filePath", "", "path to the directory that contains the emails")
var cpuProfiling = flag.Bool("cpuProfiling", false, "enable cpu profiling")
var memProfiling = flag.Bool("memProfiling", false, "enable memory profiling")
var chunks = flag.Int("chunks", 50000, "enable memory profiling")
func main() {

	// Parse the command-line flags
	flag.Parse()

	// Check if filepath is provided
	if *filePath == "" {
		fmt.Println("Error: filepath not provided")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *cpuProfiling {
		f, err := os.Create("cpu.prof")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	fmt.Printf("Parsing emails found in dir %s and its subdirs\n", *filePath)
	err := parser.ParseDir(*filePath, *chunks)

	if err != nil {
		log.Fatal("could not index the data ", err)
	}

	if *memProfiling {
		f, err := os.Create("mem.prof")
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)

		}
	}

}
