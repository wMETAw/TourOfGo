package main

import "fmt"

type IPAdder [4]byte

func (ip IPAdder) String() string {
	var s string
	for _, i := range ip {
		s += fmt.Sprint(i) + "."
	}
	return s[:len(s)-1]
}

func main() {
	hosts := map[string]IPAdder{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v:%v\n", name, ip)
	}
}
