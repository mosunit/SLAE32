/*
Tool:		AES Decrypter in Go
Author:		Mohit Suyal (@mosunit)
Student ID:	PA-16521
Blog:		https://mosunit.com
*/

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func main() {

	// define secret key as slice of byte
	// secret is same as used in encrypter shellcode
	secret := []byte("iamsecret1234569")
	fmt.Printf("The symmetric key to decrypt the shellcode is: iamsecret1234567\n\n")

	// encrypted shellcode - output of aes_shellcode_encrypter.go
	e := []byte{0x16, 0xbf, 0x7f, 0xa4, 0xda, 0x94, 0x41, 0x5b, 0x27, 0x47, 0x53, 0xbb, 0x9b, 0xb1, 0x32, 0x82, 0x60, 0x1d, 0x61, 0xcf, 0x5b, 0x05, 0x7e, 0x25, 0x1e, 0xbc, 0xec, 0x18, 0xfa, 0xe8, 0x18, 0x3e, 0xc9, 0x3e, 0xf8, 0x03, 0x2d, 0x30, 0xe0, 0xfa, 0x29, 0xb8, 0xb4, 0xc0, 0x4f, 0xd0, 0xa4, 0x32, 0x4f, 0x30, 0xb1, 0x1f, 0xbd}

	// encrypted shellcode converted to string - to be fed in decrypt function
	encrypted := fmt.Sprintf((base64.URLEncoding.EncodeToString(e)))

	decrypted, err := decrypt(encrypted, secret)
	if err != nil {
		panic(fmt.Sprintf("unable to decrypt the data: %v", err))
	}

	fmt.Println("Shellcode has been decrypted successfully. The decrypted shellcode is:")
	for i := range decrypted {
		// check the index value to match the last element of the slice
		// this if statement is true for all index values except the last
		if i != len(decrypted)-1 {
			// Check if the hex coversion of slice element will be less than 2 digits; append an additional "0", if true
			if decrypted[i] < 16 {
				fmt.Printf("\\x0%x", decrypted[i])
			} else {
				fmt.Printf("\\x%x", decrypted[i])
			}
		} else {
			if decrypted[i] < 16 {
				fmt.Printf("\\x0%x", decrypted[i])
			} else {
				fmt.Printf("\\x%x", decrypted[i])
			}
		}
	}

}

func decrypt(encodedData string, secret []byte) ([]byte, error) {
	encryptData, err := base64.URLEncoding.DecodeString(encodedData)
	if err != nil {
		return []byte{1}, err
	}

	cipherBlock, err := aes.NewCipher(secret)
	if err != nil {
		return []byte{1}, err
	}

	aead, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return []byte{1}, err
	}

	nonceSize := aead.NonceSize()
	if len(encryptData) < nonceSize {
		return []byte{1}, err
	}

	nonce, cipherText := encryptData[:nonceSize], encryptData[nonceSize:]
	plainData, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return []byte{1}, err
	}

	return plainData, nil
}
