section .data
    msg: db "Please enter your usr_name", 0
    fmt_in: dq "%s", 0
    fmt_out: db "%s", 10, 0 
    string1 db "One for "
    len1: equ $-string1
    string2 db ", one for me.", 0
    len2: equ $-string2

section .bss
    usr_name: resb 10
    output: resb 30

section .text
    global main
    extern scanf
    extern printf

_strlen:
    push rbx
    xor rbx,rbx
_loop:
    cmp byte[rsi], 0
    jz _end
    inc rbx
    inc rsi
    jmp _loop
_end:
    mov rax,rbx
    pop rbx
    ret

main:
    push rbp

    mov rdi, fmt_out
    mov rsi, msg 
    mov rax, 0
    call printf

    mov rax, 0
    mov rsi, usr_name
    mov rdi, fmt_in
    call scanf

    mov rsi, usr_name
    call _strlen

    mov rsi,string1
    mov rdi, output
    mov rcx, len1
    rep movsb

    mov rsi,usr_name
    mov rcx, rax
    rep movsb

    mov rsi,string2
    mov rcx, len2
    rep movsb

    mov rdi, fmt_out
    mov rsi, output 
    mov rax, 0
    call printf

    mov rax, 0
    pop rbp
    ret
