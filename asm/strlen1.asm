section .data
  msg: dq "Hello World", 10, 0
section .bss
section .text
  global main
  extern printf

_strlen:
  push rcx,
  xor rcx, rcx

_strlen_next:
  cmp [rdi], byte 0 ; null byte yet?
  jz _strlen_null   ; yes, get out

  inc rcx	    ; char is ok, count it
  inc rdi	    ; move to next char
  jmp _strlen_next

_strlen_null:
  mov rax, rcx	    ; rcx = the length (put in rax)
  pop rcx	    ; restore rcx
  ret		    ; get out

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

