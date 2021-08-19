section .data
    msg: db "Please enter your name", 0
    fmt_in: dq "%s", 0
    fmt_out: db "%s", 10, 0 
    string1 db "One for "
    string2 db ", one for me."

section .bss
    name: resb 20
    output: resb 50

section .text
    global main
    extern scanf
    extern printf

main:
    push rbp

    mov rdi, fmt_out
    mov rsi, msg 
    mov rax, 0
    call printf

    mov rax, 0
    mov rsi, name
    mov rdi, fmt_in
    call scanf

		mov rsi,string1
		mov rdi, output
		mov rcx, 8
		rep movsb

		mov rsi,name
		mov rdi, output
		lea rdi, [rdi + 8]
		mov rcx, 20
		rep movsb

		mov rsi,string2
		mov rdi, output
		lea rdi, [rdi + 13]
		mov rcx, 20
		rep movsb

    mov rdi, fmt_out
    mov rsi, output 
    mov rax, 0
    call printf

    mov rax, 0
    pop rbp
    ret
