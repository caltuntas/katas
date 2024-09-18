global fib

fib:
    push rbp
    mov rbp, rsp
    sub rsp, 16             ; Reserve space on the stack
    mov [rbp - 8], rdi      ; Store n in [rbp - 8]

    cmp dword [rbp - 8], 0x0
    je return_0             ; If n == 0, return 0

    cmp dword [rbp - 8], 0x1
    je return_1             ; If n == 1, return 1

    mov rdi, [rbp - 8]      ; Load n
    sub rdi, 1              ; n - 1
    call fib                ; fib(n - 1)
    mov [rbp - 16], rax     ; Save result of fib(n - 1)

    mov rdi, [rbp - 8]      ; Load n
    sub rdi, 2              ; n - 2
    call fib                ; fib(n - 2)

    add rax, [rbp - 16]     ; fib(n - 1) + fib(n - 2)
    jmp return              ; Return result

return_0:
    mov rax, 0              ; Return 0
    jmp return

return_1:
    mov rax, 1              ; Return 1
    jmp return

return:
    mov rsp, rbp            ; Restore stack pointer
    pop rbp                 ; Restore base pointer
    ret
