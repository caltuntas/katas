section .data
  msg: dq "Hello World", 10, 0
section .bss
section .text
  global main
  extern printf

_strlen:
  push rbx		;save any registers that
  push rcx 		;we will trash in there
  mov rbx, rdi		;rbx = rdi
  
  xor al,al		; the byte that the scan will compare to is zero

  mov rcx, 0xffffffff	; the max number of bytes is 4gb

  repne scasb		;while [rdi] != al, keep scanning

  sub rdi, rbx		; length = dist2 - dist1
  mov rax, rdi		; rax now holds our length
  
  pop rcx		; restore the saved registers
  pop rbx

  ret


main: 
  mov rdi, msg
  call _strlen
  mov rdx, rax

  mov rax, 1
  mov rdi, 1
  mov rsi, msg
  mov rdx, 11
  syscall

  mov rax, 60
  mov rdi, 0
  syscall

