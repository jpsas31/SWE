package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jpsas31/SWE/indexer/chiAPI"
)

var addr = flag.String("address", "localhost:4040", "address to host the server")
var profiling = flag.Bool("profiling", false, "enable cpu profiling")
var help = flag.Bool("help", false, "Explanation of the cmd options")

func main() {
	// Parse the command-line flags
	flag.Parse()
	// Check if filepath is provided
	if *help {
		fmt.Println("fgadf")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("email visualizer backend running in ",*addr)
	chiAPI.InitServer(*addr, *profiling)

}
