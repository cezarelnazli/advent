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

day3: day3.S
	$(AS) $< -o day3.o
	ld -lc day3.o -o $@
	chmod +x $@

day4: day4.S
	$(AS) $< -o day4.o
	ld -lc day4.o -o $@
	chmod +x $@

day5: day5.S
	$(AS) $< -o day5.o
	ld -lc day5.o -o $@
	chmod +x $@

clean:
	rm -f a.out *.o
	rm -f day1b day2a day2b day3 day4 day5
