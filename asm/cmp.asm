;+--------------+----+----+
;| cmp dst, src | ZF | CF |
;+--------------+----+----+
;| dst = src    |  1 |  0 |
;| dst < src    |  0 |  1 |
;| dst > src    |  0 |  0 |
;+--------------+----+----+


section .data
    msg: db 0,"Test", 0

section .text
    global main

main:
    push rbp

    ;xor rbx, rbx
    ;xor rax, rax
    xor rcx, rcx
    xor rdx, rdx

    mov rdi, msg
    cmp byte[rdi], 0

    mov rax,5
    mov rbx,8
    cmp rax,rbx 

    mov rax,8
    mov rbx,5
    cmp rax,rbx

    mov rax,5
    mov rbx,rax
    cmp rax,rbx


    pop rbp
    ret
