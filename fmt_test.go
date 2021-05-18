package fmt_test

import (
	"os"

	"4d63.com/fmt"
)

func ExamplePrintfln() {
	fmt.Printfln("Hello %s!", "world")
	fmt.Printfln("G'day %s!", "world")
	// Output:
	// Hello world!
	// G'day world!
}

func ExampleFprintfln() {
	fmt.Fprintfln(os.Stdout, "Hello %s!", "world")
	fmt.Fprintfln(os.Stdout, "G'day %s!", "world")
	// Output:
	// Hello world!
	// G'day world!
}

func ExampleSprintfln() {
	fmt.Print(fmt.Sprintfln("Hello %s!", "world"))
	fmt.Print(fmt.Sprintfln("G'day %s!", "world"))
	// Output:
	// Hello world!
	// G'day world!
}
