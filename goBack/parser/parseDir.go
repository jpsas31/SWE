package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/jpsas31/SWE/indexer/zincSearchAPIClient"
)

const INDEXINFO = "{ \"index\" : {\"_index\" : \"Emails\" }}\n"

// ParseDir parses all files in the given directory and indexes them in chunks
func ParseDir(dir string, chunks int) error {
	var buffer bytes.Buffer
	var counter int

	err := filepath.WalkDir(dir, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if !d.IsDir() {
			mailEntry, err := newEmailEntry(s)
			if err != nil {
				return err
			}
			if mailEntry == nil {
				return nil
			}

			jsonData, err := json.Marshal(mailEntry)
			if err != nil {
				return err
			}

			entry := []byte(INDEXINFO + string(jsonData) + "\n")
			buffer.Write(entry)

			counter++
			if counter%chunks == 0 {
				if err := zincSearchAPIClient.BulkIndex(buffer.Bytes()); err != nil {
					return err
				}
				counter = 0
				buffer.Reset()
			}
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking directory: %v", err)
	}

	if buffer.Len() > 0 {
		if err := zincSearchAPIClient.BulkIndex(buffer.Bytes()); err != nil {
			return err
		}
	}

	return nil
}
