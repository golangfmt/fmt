package fmt

import (
	"fmt"
	"io"
)

type (
	Formatter  = fmt.Formatter
	GoStringer = fmt.GoStringer
	ScanState  = fmt.ScanState
	Scanner    = fmt.Scanner
	State      = fmt.State
	Stringer   = fmt.Stringer
)

// Printfln is the same as fmt.Printf, but appends a new line.
func Printfln(format string, a ...interface{}) (n int, err error) {
	return Printf(format+"\n", a...)
}

// Fprintfln is the same as fmt.Fprintf, but appends a new line.
func Fprintfln(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return Fprintf(w, format+"\n", a...)
}

// Sprintf is the same as fmt.Printf, but appends a new line.
func Sprintfln(format string, a ...interface{}) string {
	return Sprintf(format+"\n", a...)
}

// Errorf is the same as fmt.Errorf.
func Errorf(format string, a ...interface{}) error { return fmt.Errorf(format, a...) }

// Fprint is the same as fmt.Fprint.
func Fprint(w io.Writer, a ...interface{}) (n int, err error) { return fmt.Fprint(w, a...) }

// Fprintf is the same as fmt.Fprintf.
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, format, a...)
}

// Fprintln is the same as fmt.Fprintln.
func Fprintln(w io.Writer, a ...interface{}) (n int, err error) { return fmt.Fprintln(w, a...) }

// Fscan is the same as fmt.Fscan.
func Fscan(r io.Reader, a ...interface{}) (n int, err error) { return fmt.Fscan(r, a...) }

// Fscanf is the same as fmt.Fscanf.
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error) {
	return fmt.Fscanf(r, format, a...)
}

// Fscanln is the same as fmt.Fscanln.
func Fscanln(r io.Reader, a ...interface{}) (n int, err error) { return fmt.Fscanln(r, a...) }

// Print is the same as fmt.Print.
func Print(a ...interface{}) (n int, err error) { return fmt.Print(a...) }

// Printf is the same as fmt.Printf.
func Printf(format string, a ...interface{}) (n int, err error) { return fmt.Printf(format, a...) }

// Println is the same as fmt.Println.
func Println(a ...interface{}) (n int, err error) { return fmt.Println(a...) }

// Scan is the same as fmt.Scan.
func Scan(a ...interface{}) (n int, err error) { return fmt.Scan(a...) }

// Scanf is the same as fmt.Scanf.
func Scanf(format string, a ...interface{}) (n int, err error) { return fmt.Scanf(format, a...) }

// Scanln is the same as fmt.Scanln.
func Scanln(a ...interface{}) (n int, err error) { return fmt.Scanln(a...) }

// Sprint is the same as fmt.Sprint.
func Sprint(a ...interface{}) string { return fmt.Sprint(a...) }

// Sprintf is the same as fmt.Sprintf.
func Sprintf(format string, a ...interface{}) string { return fmt.Sprintf(format, a...) }

// Sprintln is the same as fmt.Sprintln.
func Sprintln(a ...interface{}) string { return fmt.Sprintln(a...) }

// Sscan is the same as fmt.Sscan.
func Sscan(str string, a ...interface{}) (n int, err error) { return fmt.Sscan(str, a...) }

// Sscanf is the same as fmt.Sscanf.
func Sscanf(str string, format string, a ...interface{}) (n int, err error) {
	return fmt.Sscanf(str, format, a...)
}

// Sscanln is the same as fmt.Sscanln.
func Sscanln(str string, a ...interface{}) (n int, err error) { return fmt.Sscanln(str, a...) }
