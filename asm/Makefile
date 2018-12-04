day1b: day1b.S
	$(AS) $< -o day1b.o
	ld -lc day1b.o -o $@
	chmod +x $@

day2a: day2a.S
	$(AS) $< -o day2a.o
	ld -lc day2a.o -o $@
	chmod +x $@

day2b: day2b.S
	$(AS) $< -o day2b.o
	ld -lc day2b.o -o $@
	chmod +x $@

clean:
	rm -f a.out *.o
	rm -f day1b day2a day2b
