.global _main
fmt_str: .asciz "%s\n"
print_int: .asciz "%lld\n"

_main:
	push %rbp
	mov %rsp, %rbp

	sub $8000, %rsp		# strings
	mov $0, %rbx

read_one:
	lea fmt_str(%rip), %rdi
	mov %rbx, %rsi
	shl $5, %rsi
	add %rsp, %rsi
	call _scanf

	cmp $1, %rax
	jne read_done

	inc %rbx
	jmp read_one

read_done:
	mov $0, %rax
outer_loop:
	mov %rax, %rcx
	add $1, %rcx

inner_loop:
	cmp %rcx, %rbx
	je done_inner_loop

	mov %rax, %rsi
	shl $5, %rsi
	add %rsp, %rsi

	mov %rcx, %rdi
	shl $5, %rdi
	add %rsp, %rdi

	mov $0, %rdx

cmp_loop:
	cmp $2, %rdx
	je done_cmp_loop

	cmpb $0, (%rsi)
	je done_cmp_loop

	push %rax
	movb (%rdi), %al
	cmpb (%rsi), %al
	pop %rax
	je inc_cmp_loop
	inc %rdx
inc_cmp_loop:
	inc %rsi
	inc %rdi
	jmp cmp_loop

done_cmp_loop:
	cmp $1, %rdx
	je done
	inc %rcx
	jmp inner_loop

done_inner_loop:
	inc %rax
	jmp outer_loop

done:
	mov %rax, %rsi
	shl $5, %rsi
	add %rsp, %rsi

	mov %rcx, %rdi
	shl $5, %rdi
	add %rsp, %rdi

done_loop:
	cmpb $0, (%rsi)
	je epilogue

	push %rax
	mov (%rdi), %al
	cmpb (%rsi), %al
	pop %rax
	jne done_loop_upd

br1:
	push %rax
	push %rdi
	push %rsi
	push %rsi
	mov $0, %rax
	movb (%rdi), %al
	mov %rax, %rdi
	call _putchar
	pop %rsi
	pop %rsi
	pop %rdi
	pop %rax

done_loop_upd:
	inc %rsi
	inc %rdi
	jmp done_loop

epilogue:
	add $8000, %rsp

	mov $0, %rax
	leave
	ret
