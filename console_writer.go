package logging

import "fmt"

// ConsoleWriter Backend that uses the fmt package for printing to standard output
type ConsoleWriter struct{}

func (backend ConsoleWriter) write(s string) {
	fmt.Print(s)
}
