section .data
  msg: dq "The sorted array is: ", 10, 0
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
	mov RAX, 0
	mov RBX, 0
	mov RCX, 0
	mov RDI, msg
	call printf
	mov RAX, 0
	mov RBX, 0
	mov RCX, 0		
  ret
