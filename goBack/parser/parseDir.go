package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"sync"

	"github.com/jpsas31/SWE/indexer/zincSearchAPIClient"
)

const INDEXINFO = "{ \"index\" : {\"_index\" : \"Emails\" }}\n"

// ParseDir parses all files in the given directory and indexes them in chunks
func ParseDir(dir string, chunks int) error {
	var buffer bytes.Buffer
	var counter int
	var waitGroup sync.WaitGroup //used to wait for the multiple goroutines to finish before continuing
	var mutex sync.Mutex //mutual exclusion used to 

	err := filepath.WalkDir(dir, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if !d.IsDir() {
			waitGroup.Add(1) //adds a new goroutine to the waiting group
			go func(file string) {
				defer waitGroup.Done()// substracts a goroutine from the waiting group. this is done when the func is done thanks to defer

				mailEntry, err := newEmailEntry(file)
				if err != nil {
					
					return
				}
				if mailEntry == nil {
					return
				}

				jsonData, err := json.Marshal(mailEntry)
				if err != nil {
					
					return
				}

				entry := []byte(INDEXINFO + string(jsonData) + "\n")

				mutex.Lock() // locks the shared resources (buffer and counter) so that only one go routine access them at a time
				defer mutex.Unlock() // when the function is done it realeases the resources for another go routine

				buffer.Write(entry)
				counter++
				if counter%chunks == 0 {
					if err := zincSearchAPIClient.BulkIndex(buffer.Bytes()); err != nil {
						return
					}
					counter = 0
					buffer.Reset()
				}
			}(s)
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking directory: %v", err)
	}

	waitGroup.Wait()

	if buffer.Len() > 0 {
		if err := zincSearchAPIClient.BulkIndex(buffer.Bytes()); err != nil {
			return err
		}
	}

	return nil
}
