// +build js

package main

import (
	"io"

	"honnef.co/go/js/dom"
)

var document = dom.GetWindow().Document()

func init() {
	stdin = NewReader(document.GetElementByID("input").(*dom.HTMLInputElement))
	stdout = NewWriter(document.GetElementByID("output").(*dom.HTMLPreElement))
	stderr = NewWriter(document.GetElementByID("output").(*dom.HTMLPreElement))
}

// end OMIT

var _ = io.Reader
