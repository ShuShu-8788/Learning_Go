package main

import (
	"fmt"

	"rsc.io/quote"

	"greetings"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
	fmt.Println("ShuShu-Day-2")

	//Get greeting module and print it.
	message := greetings.Hello("Hello Go Gladys <Passed Arguments>")

	fmt.Println(message)
}
