
section .data
    msg: db "Test", 0

section .text
    global main

main:
    push rbp

    xor rbx, rbx
    xor rax, rax
    xor rcx, rcx
    mov rdi, msg
    lea rsi, [msg] 

    movsx rbx, byte [msg]
    mov cl, byte [msg]
    mov al, [msg]

    pop rbp
    ret
