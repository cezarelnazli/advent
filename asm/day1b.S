.bss:
	.lcomm seen, 200000
	.lcomm nums, 800000

.text:
	.global _main
	scanf_str: .asciz " %ld"
	printf_str: .asciz "%d\n"

_main:
	push %rbp
	mov %rsp, %rbp

	lea nums(%rip), %rbx
	mov $0, %r12

read_one:
	lea scanf_str(%rip), %rdi

	lea 0(%rbx, %r12, 4), %rsi

	call _scanf

	cmp $1, %rax
	jne read_done

	inc %r12

	jmp read_one

read_done:
	mov $0, %rax
	lea seen(%rip), %rdx
	mov $0, %rcx

loop:
	add (%rbx, %rax, 4), %ecx

	push %rax
	mov 100000(%edx, %ecx), %al

	cmpb $1, %al
	je print_res

	pop %rax

	movb $1, 100000(%edx, %ecx)
	inc %rax

	cmp %rax, %r12
	jne dont_reset
	mov $0, %rax
dont_reset:
	jmp loop

print_res:
	pop %rax
	lea printf_str(%rip), %rdi
	mov %rcx, %rsi
	call _printf

	mov $0, %rax

	leave
	ret
