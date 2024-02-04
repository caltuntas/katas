//clang -arch x86_64 hello_macos_intel.asm -o hellomacos-intel
//.intel_syntax
.intel_syntax noprefix
.data
msg:
    .ascii "Hello, world!\n"

.text
    .globl _main

_main:
    sub  rsp, 8                  
    mov  rax, 0
    lea  rdi, [rip + msg]      
    call _printf                 
    add  rsp, 8                  
    ret
