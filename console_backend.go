package logging

import "fmt"

// ConsoleBackend Backend that uses the fmt package for printing to standard output
type ConsoleBackend struct{}

func (backend ConsoleBackend) log(s string) {
	fmt.Print(s)
}
