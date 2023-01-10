package handlers

import (
	"log"
	"net/http"
)
// Goodbye is a simple handler
type Goodbye struct {
	l *log.Logger
}
// NewHGoodbye creates a new goodbye handler with the given logge
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}
// ServeHTTP implements the go http.Handler interface
// https://golang.org/pkg/net/http/#Handler
func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byee"))
}
