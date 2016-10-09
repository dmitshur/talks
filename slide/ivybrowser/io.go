// +build ignore

package main

import (
	"io"

	"honnef.co/go/js/dom"
)

// NewReader start OMIT

// NewReader takes an <input> element and makes an io.Reader out of it.
func NewReader(input *dom.HTMLInputElement) io.Reader {
	r := &reader{
		in: make(chan []byte, 8),
	}
	input.AddEventListener("keydown", false, func(event dom.Event) {
		ke := event.(*dom.KeyboardEvent)
		if ke.KeyCode == '\r' {
			r.in <- []byte(input.Value + "\n") // HL1
			input.Value = ""                   // HL1
			ke.PreventDefault()
		}
	})
	return r
}

// NewReader end OMIT

// reader start OMIT

type reader struct {
	pending []byte
	in      chan []byte // This channel is never closed, no need to detect it and return io.EOF.
}

func (r *reader) Read(p []byte) (n int, err error) {
	if len(r.pending) == 0 {
		r.pending = <-r.in // HL1
	}
	n = copy(p, r.pending)
	r.pending = r.pending[n:]
	return n, nil
}

// reader end OMIT

// NewWriter start OMIT

// NewWriter takes a <pre> element and makes an io.Writer out of it.
func NewWriter(pre *dom.HTMLPreElement) io.Writer {
	return &writer{pre: pre}
}

type writer struct {
	pre *dom.HTMLPreElement
}

func (w *writer) Write(p []byte) (n int, err error) {
	w.pre.SetTextContent(w.pre.TextContent() + string(p)) // HL1
	return len(p), nil
}

// NewWriter end OMIT
