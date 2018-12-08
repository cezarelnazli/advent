.bss:
	.lcomm claimed, 125000
	.lcomm accounted, 125000

.text:
	.global _main
	scanf_str: .asciz "#%lld @ %lld,%lld: %lldx%lld\n"
	printf_str: .asciz "%d\n"

_main:
	push %rbp
	mov %rsp, %rbp

	sub $0x80, %rsp
	mov $0, %rbx
	mov $0, %rsi
	movq $0, 0x50(%rsp) # prev ID
	movq $0, 0x60(%rsp) # is part two?
	movq $0, 0x70(%rsp) # part two solution

read_one:
	lea scanf_str(%rip), %rdi
	lea 0x40(%rsp), %rsi
	lea (%rsp), %rdx
	lea 0x10(%rsp), %rcx
	lea 0x20(%rsp), %r8
	lea 0x30(%rsp), %r9
	call _scanf

	# XXX: Part 2: comment these two lines and give input twice
	# Example:
	#
	# #1 @ 1,3: 4x4
	# #2 @ 3,1: 4x4
	# #3 @ 5,5: 2x2
	# #1 @ 1,3: 4x4
	# #2 @ 3,1: 4x4
	# #3 @ 5,5: 2x2
	#
	cmp $5, %rax
	jne done

	cmpq $0, 0x70(%rsp)
	jne done_part_two

	mov $0, %r14
	mov (%rsp), %r11
	mov 0x10(%rsp), %r12
	mov 0x20(%rsp), %rcx
	mov 0x30(%rsp), %r13
	mov 0x40(%rsp), %rsi

	cmp 0x50(%rsp), %rsi
	ja part_one

	movq $1, 0x60(%rsp)

part_one:
	mov %rsi, 0x50(%rsp)

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
	jnz col_loop_upd_zero

	inc %rbx
	or %al, %dl
	movb %dl, (%rsi)

	push %r10
	mov -0x40(%rbp), %r10
	mov %r10, -0x10(%rbp)
	pop %r10

	cmpq $1, -0x20(%rbp)
	je col_loop_upd
	jmp col_loop_upd_zero

not_claimed:
	or %al, %dl
	movb %dl, (%rdi)
	jmp col_loop_upd

col_loop_upd_zero:
	pop %rcx
	inc %r15
	movq $0, 0x70(%rsp)
	cmpq $1, -0x20(%rbp)
fb:
	je read_one
	jmp col_loop

col_loop_upd:
	pop %rcx
	inc %r15
	jmp col_loop

row_loop_upd:
	inc %r14
	jmp row_loop

done_part_two:
	mov 0x70(%rsp), %rbx

done:
	lea printf_str(%rip), %rdi
	mov %rbx, %rsi
	call _printf

	add $0x80, %rsp
	mov $0, %rax
	leave
	ret
