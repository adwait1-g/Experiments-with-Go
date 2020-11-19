package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// String() function
func (ip IPAddr) String() string {
	if len(ip) != 4 {
		return ""
	}

	var result string
	for _, number := range ip {
		result += strconv.Itoa((int(number)))
		result += "."
	}

	result = result[0 : len(result)-1]
	return result
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
