default rel            ; make [rel format] the default, you always want this.
extern printf, exit    ; NASM requires declarations of external symbols, unlike GAS
section .rodata
    format db "%#x", 10, 0   ; C 0-terminated string: "%#x\n" 
section .text
global main
main:
    sub   rsp, 8             ; re-align the stack to 16 before calling another function

    ; Call printf.
    mov   esi, 0x12345678    ; "%x" takes a 32-bit unsigned int
    lea   rdi, [rel format]
    xor   eax, eax           ; AL=0  no FP args in XMM regs
    call  printf

    ; Return from main.
    xor   eax, eax
    add   rsp, 8
    ret
