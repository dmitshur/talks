// +build js

package main

import (
	"io"

	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document()

func init() {
	stdin = NewReader(document.GetElementByID("input").(*dom.HTMLInputElement)) // HL
	stdout = NewWriter(document.GetElementByID("output").(*dom.HTMLPreElement)) // HL
	stderr = NewWriter(document.GetElementByID("output").(*dom.HTMLPreElement)) // HL
}

// end OMIT

var _ = io.Reader
