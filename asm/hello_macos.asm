# clang -arch x86_64 hello_macos.asm -o hellomacos
.cstring
_hello: .asciz "Hello, world\n"

.text
.globl _main
_main:
    sub  $8, %rsp           // align rsp to 16B boundary
    mov  $0, %rax
    lea  _hello(%rip), %rdi // load address of format string
    call _printf            // call printf
    add  $8, %rsp           // restore rsp
    ret
