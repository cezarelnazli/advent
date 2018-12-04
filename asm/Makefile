day1b: day1b.S
	$(AS) $< -o day1b.o
	ld -lc day1b.o -o $@
	chmod +x $@

clean:
	rm -f a.out day1b *.o
