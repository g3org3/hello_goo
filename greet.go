package hello_goo

import "fmt"

// Hello returns a friendly greeting for the given name.
func Hello(name string) string {
	return fmt.Sprintf("Hey, %s!", capitalize(name))
}

// Formal returns a formal greeting for the given name.
func Formal(name string) string {
	return fmt.Sprintf("Good day, %s. Welcome.", capitalize(name))
}
