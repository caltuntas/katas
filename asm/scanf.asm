section .data
    msg: db "Please enter your usr_name", 0
    fmt_in: dq "%s", 0
    fmt_out: db "%s", 10, 0 
    string1 db "One for ", 0
    string2 db ", one for me.", 0

section .bss
    usr_name: resb 10
    output: resb 30

section .text
    global main
    extern scanf
    extern printf

;rsi source string
;rdi destination string
_strcat:
    push rcx
    xor rcx,rcx

    mov rdx,rdi
    call _strlen
    lea rdi,[rdi + rax]
    mov rdx,rsi
    call _strlen
    mov rcx, rax
    rep movsb

    mov rax, rdi
    pop rcx
    ret

_strlen:
    push rbx
    xor rbx,rbx
_loop:
    cmp byte[rdx], 0
    jz _end
    inc rbx
    inc rdx
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

    mov rsi,string1
    mov rdi, output
    call _strcat

    mov rdx,usr_name
    call _strlen
    mov rsi,usr_name
    mov rcx, rax
    rep movsb

    mov rdx,string2
    call _strlen
    mov rsi,string2
    mov rcx, rax
    rep movsb

    mov rdi, fmt_out
    mov rsi, output 
    mov rax, 0
    call printf

    mov rax, 0
    pop rbp
    ret
