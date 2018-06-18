package logging

// Backend Interface that logging backends need to fulfill
type Backend interface {
	log(s string)
}
