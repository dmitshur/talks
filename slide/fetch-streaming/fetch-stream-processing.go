package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
)

// Search start OMIT

func Search(url, query string) (result string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close() // HL
	r := csv.NewReader(resp.Body) // HL
	for {
		record, err := r.Read()
		if err == io.EOF {
			return "", fmt.Errorf("could not find value after %q", query)
		}
		if err != nil {
			// handle error
		}
		// query logic ...
	}
}

// Search end OMIT

//func main() {
// main start OMIT
result, err := Search("/large.csv", "gopher")
if err != nil {
	// handle error
}
fmt.Printf("Got the result! It's %q\n", result)
// main end OMIT
//}
