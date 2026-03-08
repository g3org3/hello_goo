package hello_goo

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hey, %s!", name)
}

func Formal(name string) string {
	return fmt.Sprintf("Good day, %s. Welcome.", name)
}
