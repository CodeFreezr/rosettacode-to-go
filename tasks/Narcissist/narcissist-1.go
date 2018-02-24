package main

import "os"
import "fmt"
import "bytes"
import "io/ioutil"

func main() {
	ios := "os"
	ifmt := "fmt"
	ibytes := "bytes"
	iioutil := "io/ioutil"
	zero := "Reject"
	one := "Accept"
	x := "package main; import %q; import %q; import %q; import %q; func main() {ios := %q; ifmt := %q; ibytes := %q; iioutil := %q; zero := %q; one := %q; x := %q; s := fmt.Sprintf(x, ios, ifmt, ibytes, iioutil, ios, ifmt, ibytes, iioutil, zero, one, x); in, _ := ioutil.ReadAll(os.Stdin); if bytes.Equal(in, []byte(s)) {fmt.Println(one);} else {fmt.Println(zero);};}\n"
	s := fmt.Sprintf(x, ios, ifmt, ibytes, iioutil, ios, ifmt, ibytes, iioutil, zero, one, x)
	in, _ := ioutil.ReadAll(os.Stdin)
	if bytes.Equal(in, []byte(s)) {
		fmt.Println(one)
	} else {
		fmt.Println(zero)
	}
}

//\Narcissist\narcissist-1.go
