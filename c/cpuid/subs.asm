
section .text
	global cpuidfn

cpuidfn:
	push rbp
        mov rbp, rsp
        mov rax, rdi
        cpuid
	mov rax, rcx
        leave
        ret