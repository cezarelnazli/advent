def part_one
    all_lines = File.read('input/checksum.in')
    sum = 0

    all_lines.each_line do |l|
        numbers = l.split.map {|x| x.to_i}
        lmin = numbers[0]
        lmax = numbers[0]
        numbers.each do |n|
            if n < lmin
                lmin = n
            end

            if n > lmax
                lmax = n
            end
        end

        sum += lmax - lmin
    end

    puts sum
end

def gcd(a, b)
    if a % b == 0
        b
    else
        gcd(b, a % b)
    end
end

def part_two
    all_lines = File.read('input/checksum.in')
    sum = 0

    all_lines.each_line do |l|
        numbers = l.split.map {|x| x.to_i}
        n = numbers.length

        i = 0
        loop do
            j = 0
            loop do
                if i != j && numbers[i] >= numbers[j] && gcd(numbers[i], numbers[j]) == numbers[j]
                    sum += numbers[i] / numbers[j]
                    break
                end

                j += 1
                if j == n
                    break
                end
            end

            i += 1
            if i == n
                break
            end
        end
    end

    puts sum
end
