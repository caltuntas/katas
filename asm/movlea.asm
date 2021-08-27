
section .data
    msg: db "Test", 0

section .text
    global main

main:
    push rbp

    xor rbx, rbx
    xor rax, rax
    xor rcx, rcx
    xor rdx, rdx

    mov rdi, msg
    mov rdx, msg + 1
    lea rsi, [msg] 

    movsx rbx, byte [msg]
    mov cl, byte [msg]
    mov al, [msg]

    pop rbp
    ret
