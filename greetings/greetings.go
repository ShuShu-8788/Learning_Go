package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi , %v. Welcome! from userDefined Module", name)
	return message
}
