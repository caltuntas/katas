push rbp
mov rbp, rsp
sub rsp, 200
mov [rbp-8], rdi
cmp dword [rbp-8],0
jg _L1
mov rax, 0
jmp FUNC_END
cmp dword [rbp-8],1
jne _L2
mov rax, 1
jmp FUNC_END
mov rax, [rbp -8]
sub rax, 1
mov rdi, rax
call fib
mov [rbp-16],rax
mov rdx, [rbp -8]
sub rdx, 2
mov rdi, rdx
call fib
mov [rbp-24],rax
mov rcx, [rbp -16]
add rcx, [rbp -24]
mov rax, rcx
jmp FUNC_END
mov rsp, rbp
pop rbp
ret
