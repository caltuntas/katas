;int fib(int n) {
  ;    if (n <= 0)
    ;        return 0;
  ;    if (n == 1)
    ;        return 1;
  ;    return fib(n - 1) + fib(n - 2);
  ;}
global fib

fib:
  push rbp
  mov rbp,rsp
  push rdx

  ;check if n == 0
  cmp rdi, 0x0
  je return_0

  ;check if n == 1
  cmp rdi, 0x1
  je return_1

  ;fib(n-1)
  push rdi
  sub rdi, 0x1
  call fib
  mov rdx, rax
  pop rdi

  ;push rdi
  ;fib(n-2)
  sub rdi, 0x2
  call fib
  add rdx, rax
  ;pop rdi

  mov rax, rdx
  jmp return

return_0:
  mov rax, 0x0
  jmp return


return_1:
  mov rax, 0x1
  jmp return


return:
  pop rdx
  mov rsp, rbp
  pop rbp
  ret
