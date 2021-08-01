section .data
  msg: dq "The sorted array is: ", 10, 0
  printf_format: db "%x",10,0
section .bss
section .text
  global main
  extern printf

doit:
  push rbp ; and this push decreases the stack 
  mov rbp, rsp
  ; call pushed the return address onto the stack, so that is 8 bytes (64 bits)
  ; push rbp pushed another 8 bytes. So to access the first arg a have to 
  ; specify 8+8=16
  mov rax, [rbp+16]
  leave
  ret

main: 
  push rbp
  mov rbp, rsp
  push $18
  call doit
  leave
  sub   rsp, 8             ; re-align the stack to 16 before calling another function

  ; Call printf.
  mov   esi, 0x12345678    ; "%x" takes a 32-bit unsigned int
  lea   rdi, [rel printf_format]
  xor   eax, eax           ; AL=0  no FP args in XMM regs
  call  printf

  ; Return from main.
  xor   eax, eax
  add   rsp, 8
  ret
