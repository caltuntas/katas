extern printf
section .rodata
    fmt1 db "value of r12=%#x", 10, 0
    fmt2 db "value of r13=%#x", 10, 0
section .text
global main
main:
    sub   rsp, 8   
    mov   r12, 0x10
    mov   r13, 0x20

    xor   r12, r13
    xor   r13, r12
    xor   r12, r13

    mov   rsi, r12 
    lea   rdi, [rel fmt1]
    xor   rax, rax
    call  printf

    mov   rsi, r13 
    lea   rdi, [rel fmt2]
    xor   rax, rax
    call  printf

    xor   rax, rax
    add   rsp, 8
    ret
