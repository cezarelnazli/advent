.bss:
	.lcomm claimed, 125000
	.lcomm accounted, 125000

.text:
	.global _main
	scanf_str: .asciz "#%*d @ %lld,%lld: %lldx%lld\n"
	printf_dbg: .asciz "%d %d %d %d\n"
	printf_str: .asciz "%d\n"

_main:
	push %rbp
	mov %rsp, %rbp

	sub $0x40, %rsp
	mov $0, %rbx

read_one:
	lea scanf_str(%rip), %rdi
	lea (%rsp), %rsi
	lea 0x10(%rsp), %rdx
	lea 0x20(%rsp), %rcx
	lea 0x30(%rsp), %r8
	call _scanf

	cmp $4, %rax
	jne done

	mov $0, %r14
	mov (%rsp), %r11
	mov 0x10(%rsp), %r12
	mov 0x20(%rsp), %rcx
	mov 0x30(%rsp), %r13

row_loop:
	mov $0, %r15
	cmp %r14, %rcx
	je read_one
col_loop:
	cmp %r15, %r13
	je row_loop_upd

	push %rcx

	lea (%r11, %r14), %rax
	mov $1000, %r9
	imul %r9, %rax
	add %r12, %rax
	add %r15, %rax
	mov $0, %rdx
	mov $8, %r9
	div %r9

	lea claimed(%rip), %rdi
	add %rax, %rdi
	mov %rax, %r10
	movb (%rdi), %al

	mov %dl, %cl
	mov $1, %dl
	shlb %cl, %dl

	test %al, %dl
	jz not_claimed

	lea accounted(%rip), %rsi
	add %r10, %rsi
	movb (%rsi), %al

	test %al, %dl
	jnz col_loop_upd

	inc %rbx
	or %al, %dl
	movb %dl, (%rsi)
	jmp col_loop_upd

not_claimed:
	or %al, %dl
	movb %dl, (%rdi)

col_loop_upd:
	pop %rcx
	inc %r15
	jmp col_loop

row_loop_upd:
	inc %r14
	jmp row_loop

done:
	lea printf_str(%rip), %rdi
	mov %rbx, %rsi
	call _printf

	add $0x40, %rsp
	mov $0, %rax
	leave
	ret
