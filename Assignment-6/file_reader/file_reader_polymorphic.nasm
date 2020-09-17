global _start

_start:
	xor ebx, ebx		; EBX XORed to NULL
	push ebx		; push NULL onto stack

	; push 	/etc///netconfig on the stack
	push 0x6769666e		; gifn : 6769666e
	push 0x6f637465		; octe : 6f637465
	push 0x6e2f2f2f		; n/// : 6e2f2f2f
	push 0x6374652f		; cte/ : 6374652f
	mov ebx, esp

	xor eax, eax 		; EAX XOred to NULL
	add al, 5 		; Increment AL by 5; EAX - 0x5
	xor ecx, ecx
	int 0x80

	mov esi, eax
	jmp read

exit:
	xor eax, eax 		; EAX XOred to NULL
	add al, 1 		; Increment AL by 1; EAX - 0x1
	xor ebx, ebx
	int 0x80

read:
	mov ebx, esi
	xor eax, eax		; EAX XOred to NULL
	add al, 3		; Increment AL by 3; EAX - 0x3
	sub esp, 1
	lea ecx,[esp]
	xor edx, edx		; EDX XOred to NULL
	add dl, 1		; Increment DL by 1, EDX - 0x1
	int 0x80

	xor ebx, ebx
	cmp ebx, eax
	je exit

	xor eax, eax		; EAX XOred to NULL
	xor edx, edx		; EDX XOred to NULL
	add al, 4
	add bl, 1
	add dl, 1
	int 0x80

	add esp, 1
	jmp read
