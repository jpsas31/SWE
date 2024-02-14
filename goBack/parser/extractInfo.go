package parser

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// checkError checks for an error and handles it.
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// sumMapValuesLen calculates the sum of lengths of values in a map.
func sumMapValuesLen(m map[string]string) int {
	sum := 0
	for _, v := range m {
		sum += len(v)
	}
	return sum
}

// newEmailEntry reads an email file and returns its parsed content.
func newEmailEntry(filepath string) (map[string]string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	headerEndPos := bytes.Index(data, []byte("X-FileName"))
	headerEndPos = headerEndPos + bytes.Index(data[headerEndPos:], []byte("\n"))
	message := string(data[headerEndPos:])
	header := string(data[:headerEndPos])
	entry := make(map[string]string)
	parseHeader(header, &entry)
	entry["Message"] = message
	if sumMapValuesLen(entry) > 1000000 {
		fmt.Println("Email ", filepath, " is too long and can't be indexed")
		return nil, nil
	}

	return entry, nil
}

// parseHeader parses the email header and populates the entry map.
func parseHeader(header string, entry *map[string]string) {
	headerEntries := strings.Split(header, "\n")
	var lastField string
	for _, line := range headerEntries {
		pair := strings.SplitN(strings.Trim(line, " "), ":", 2)
		if len(pair) == 2 {
			pair[1] = strings.Trim(pair[1], " ")
		} else {
			pair = append(pair, "")
		}

		switch pair[0] {
		case "Message-ID", "Date", "From", "To",
			"Subject", "Mime-Version", "Content-Type",
			"Content-Transfer-Encoding", "X-From",
			"X-To", "X-cc", "X-bcc", "X-Folder",
			"X-Origin", "X-FileName", "Cc", "Bcc":
			{
				key := strings.TrimSpace(pair[0])
				val := strings.TrimSpace(pair[1])
				if len(val) != 0 {
					(*entry)[key] = val
					lastField = key
				}
			}
		default:
			(*entry)[lastField] += "\n" + pair[0]
		}
	}
}
