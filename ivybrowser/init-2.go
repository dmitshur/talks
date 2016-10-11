// +build ignore

package main

import (
	"io"

	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document()

// start OMIT

func init() {
	stdin = NewReader(document.GetElementByID("input").(*dom.HTMLInputElement))
	stdout = NewWriter(document.GetElementByID("output").(*dom.HTMLPreElement))
	stderr = NewWriter(document.GetElementByID("output").(*dom.HTMLPreElement))

	// Send a copy of stdin to stdout (like in most terminals). // HL
	stdin = io.TeeReader(stdin, stdout) // HL
}
