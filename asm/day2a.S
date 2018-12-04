.global _main
scanf_s: .asciz "%s\n"
printf_s: .asciz "%d\n"

_main:
	push %rbp
	mov %rsp, %rbp

	# string, count, 2's, 3's
	sub $64, %rsp
	sub $32, %rsp
	sub $8, %rsp
	sub $8, %rsp

	movq $0, (%rsp)
	movq $0, 8(%rsp)

read_one:
	lea scanf_s(%rip), %rdi
	lea 48(%rsp), %rsi
	call _scanf

	cmp $1, %rax
	jne read_done

	mov $0, %rcx
	lea 16(%rsp), %rdx

zero_count:
	cmp $26, %rcx
	je make_count
	lea (%rdx, %rcx), %rax
	movb $0, (%rax)
	inc %rcx
	jmp zero_count

make_count:
	mov $0, %rcx
	mov $0, %rdx

make_count_loop:
	movb 48(%rsp, %rcx), %dl
	cmp $0, %dl
	je update_result
	sub $97, %dl
	incb 16(%rsp, %dl)
	inc %rcx
	jmp make_count_loop

update_result:
	mov $0, %rcx
	mov $0, %rdx
	mov $0, %rbx

upd_res_loop:
	cmp $26, %rcx
	je read_one
	mov 16(%rsp, %rcx), %al
	cmp $2, %al
	jne cmp_three
	cmp $0, %rdx
	jne next
	incq 8(%rsp)
	mov $1, %rdx
	jmp next
cmp_three:
	cmp $3, %al
	jne next
	cmp $0, %rbx
	jne next
	incq (%rsp)
	mov $1, %rbx
next:
	inc %rcx
	jmp upd_res_loop

read_done:
	mov (%rsp), %rax
	mov 8(%rsp), %rsi
	imulq %rax, %rsi

	lea printf_s(%rip), %rdi
	call _printf

	add $112, %rsp

	mov $0, %rax
	leave
	ret
