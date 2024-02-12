// package main


// import (
// 	"bufio"
// 	"bytes"
// 	"fmt"
// )



// func main() {
// 	// Create a buffer with data larger than the maximum allowed token size
// 	data := bytes.Repeat([]byte("a"), bufio.MaxScanTokenSize+100)

// 	// Create a custom reader with the data
// 	reader := bytes.NewReader(data)

	
// 	scanner := bufio.NewScanner(reader)

// 	// Scan through the input
// 	for scanner.Scan() {
// 		fmt.Println("hola")
// 		// Handle each token
// 		token := scanner.Text()
// 		fmt.Println("Token:", token)
// 	}

// 	// Check for errors
// 	if err := scanner.Err(); err != nil {

// 		fmt.Println("Error:", err)

// 	}
// }

package main

import (
	"fmt"
	"os"

	"github.com/jpsas31/SWE/indexer/parser"
)

func main() {

	fmt.Printf("Parsing emails found in dir %s and its subdirs\n", os.Args[1])
	parser.ParseDir(os.Args[1], 50000)
}
