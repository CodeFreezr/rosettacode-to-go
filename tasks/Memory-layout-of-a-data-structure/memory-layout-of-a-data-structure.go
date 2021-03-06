package main

import "fmt"

type rs232p9 uint16

const (
	CD9  rs232p9 = 1 << iota // Carrier detect
	RD9                      // Received data
	TD9                      // Transmitted data
	DTR9                     // Data terminal ready
	SG9                      // signal ground
	DSR9                     // Data set ready
	RTS9                     // Request to send
	CTS9                     // Clear to send
	RI9                      // Ring indicator
)

func main() {
	// set some nonsense bits just for example
	p := RI9 | TD9 | CD9
	fmt.Printf("Type=%T value=%#04x\n", p, p)
}

//\Memory-layout-of-a-data-structure\memory-layout-of-a-data-structure.go
