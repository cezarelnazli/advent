.bss:
    .lcomm input, 65024
    .lcomm when_asleep, 426024
    .lcomm time_asleep, 6984

.text:
    .global _main
    scanf_str: .asciz "[1518-%hhd-%hhd %hhd:%hhd] %[^\n]\n"
    printf_str: .asciz "%d\n"
    guard_str: .asciz "Guard #%d begins shift"

gcmp:
    mov $0, %rax
    mov (%edi), %eax
    sub (%esi), %eax
    ret

_main:
    push %rbp
    mov %rsp, %rbp

    sub $0x10, %rsp
    mov $0, %rbx

read_one:
    lea scanf_str(%rip), %rdi

    mov %rbx, %rsi
    shl $5, %rsi
    lea input(%rip), %rdx
    add %rdx, %rsi
    add $3, %rsi

    mov %rsi, %rdx
    dec %rdx

    mov %rdx, %rcx
    dec %rcx

    mov %rcx, %r8
    dec %r8

    mov %r8, %r9
    add $4, %r9

    call _scanf

    cmp $5, %rax
    jne read_done

    inc %rbx
    jmp read_one

read_done:
    lea input(%rip), %rdi
    mov %rbx, %rsi
    mov $32, %rdx
    lea gcmp(%rip), %rcx
    call _qsort

    mov $0, %rcx
loop_logs:
    cmp %rcx, %rbx
    je parsed_logs

    mov %rcx, %rdi
    imul $32, %rdi
    add $4, %rdi
    lea input(%rip), %rax
    add %rax, %rdi
    cmpb $71, (%rdi)
    jne test_wake

    push %rcx
    push %rdi
    lea guard_str(%rip), %rsi
    lea -0x10(%rbp), %rdx
    call _sscanf
    pop %rdi
    pop %rcx

    jmp loop_logs_inc

test_wake:
    cmpb $119, (%rdi)
    jne loop_logs_inc

    sub $4, %rdi
    mov $0, %rax
    mov $0, %rdx
    movb (%rdi), %al
    sub $32, %rdi
    movb (%rdi), %dl

    lea when_asleep(%rip), %rdi
    lea time_asleep(%rip), %rsi

    mov %rsi, %r9
    mov $0, %r10
    mov -0x10(%rbp), %r10d
    shl $1, %r10
    add %r10, %r9
    shr $1, %r10

asleep_loop:
    cmp %al, %dl
    je loop_logs_inc

    incw (%r9)

    mov %rdx, %r12
    shl $1, %r12
    mov %r10, %r11
    imul $120, %r11
    add %r12, %r11
    add %rdi, %r11
    incw (%r11)

    inc %dl
    jmp asleep_loop

loop_logs_inc:
    inc %rcx
    jmp loop_logs

parsed_logs:
    mov $0, %rax
    mov $0, %rbx
    mov $0, %rcx

guard_loop:
    cmp $3491, %rax
    je found_guard

    lea (%rsi, %rax, 2), %r9

    cmpw (%r9), %cx
    jae guard_loop_inc

    mov (%r9), %cx
    mov %rax, %rbx

guard_loop_inc:
    inc %rax
    jmp guard_loop

found_guard:
    mov $0, %rax
    mov $0, %rcx
    mov $0, %rdx

minute_loop:
    cmp $60, %rax
    je print_hash

    mov %rax, %r12
    shl $1, %r12
    mov %rbx, %r11
    imul $120, %r11
    add %r12, %r11
    add %rdi, %r11

    cmpw (%r11), %cx
    jae minute_loop_inc

    mov (%r11), %cx
    mov %rax, %rdx

minute_loop_inc:
    inc %rax
    jmp minute_loop

print_hash:
    mov %rbx, %rsi
    imul %rdx, %rsi

    lea printf_str(%rip), %rdi
    call _printf

    lea when_asleep(%rip), %rdi
    mov $0, %rax
    mov $0, %rbx
    mov $0, %rdx

max_minute_loop:
    cmp $418920, %rax
    je print_second_hash

    mov %rax, %rcx
    add %rdi, %rcx
    cmpw (%rcx), %bx

    jae max_minute_loop_inc

    mov (%rcx), %bx

    push %rax
    push %rbx

    mov %rax, %rsi
    mov $120, %rbx
    mov $0, %rdx
    div %rbx
    shr $1, %rdx
fb:
    imul %rax, %rdx
    mov %rdx, %rsi

    pop %rbx
    pop %rax

max_minute_loop_inc:
    add $2, %rax
    jmp max_minute_loop

print_second_hash:
    lea printf_str(%rip), %rdi
    call _printf

done:
    add $0x10, %rsp

    mov $0, %rax
    leave
    ret
