.bss:
    .lcomm input, 50001
    .lcomm remove_polymer, 50001

.text:
    .global _main
    printf_str: .asciz "%d\n"

react:
    push %rbp
    push %rbx
    push %r12
    push %r13
    push %r14
    push %r15
    push %r15

    mov $1, %r9
infinite_loop:
    cmp $0, %r9
    je done_react

    mov $0, %r9
    mov $0, %rax
    mov $0, %rsi

string_loop:
    movb (%rdi, %rsi), %bl
    cmpb $0, %bl
    je string_loop_done

    movb 1(%rdi, %rsi), %cl

    mov $1, %rdx
    shl $5, %rdx

    push %rbx
    push %rcx

    and %dl, %bl
    and %dl, %cl

    cmp %bl, %cl
    je string_loop_inc_pop

    pop %rcx
    pop %rbx

    not %dl

    and %dl, %bl
    and %dl, %cl

    cmp %bl, %cl
    jne string_loop_inc

    add $2, %rsi
    mov $1, %r9
    jmp string_loop

string_loop_inc_pop:
    pop %rcx
    pop %rbx

string_loop_inc:
    movb (%rdi, %rsi), %cl
    movb %cl, (%rdi, %rax)

    inc %rsi
    inc %rax

    jmp string_loop

string_loop_done:
    movb $0, (%rdi, %rax)
    jmp infinite_loop

done_react:
    call _strlen

    pop %r15
    pop %r15
    pop %r14
    pop %r13
    pop %r12
    pop %rbx
    pop %rbp

    ret

_main:
    push %rbp
    mov %rsp, %rbp

    lea input(%rip), %rdi
    call _gets

    lea remove_polymer(%rip), %rdi
    lea input(%rip), %rsi
    call _strcpy

    lea remove_polymer(%rip), %rdi

    call react

    mov %rax, %rbx

    lea printf_str(%rip), %rdi
    mov %rax, %rsi
    call _printf

    movb $65, %cl
    movb $97, %dl
remove_loop:
    cmp $92, %cl
    je remove_loop_done

    lea remove_polymer(%rip), %rdi
    lea input(%rip), %rsi

cpy_loop:
    cmpb (%rsi), %cl
    je cpy_loop_inc
    cmpb (%rsi), %dl
    je cpy_loop_inc
    cmpb $0, (%rsi)
    je cpy_loop_done

    movb (%rsi), %r9b
    movb %r9b, (%rdi)

    inc %rdi

cpy_loop_inc:
    inc %rsi
    jmp cpy_loop

cpy_loop_done:
    movb $0, (%rdi)

    lea remove_polymer(%rip), %rdi

    push %rcx
    push %rdx
    call react
    pop %rdx
    pop %rcx

    cmp %rax, %rbx
    jbe remove_loop_inc

    mov %rax, %rbx

remove_loop_inc:
    inc %cl
    inc %dl
    jmp remove_loop

remove_loop_done:
    lea printf_str(%rip), %rdi
    mov %rbx, %rsi
    call _printf

    mov $0, %rax
    leave
    ret
