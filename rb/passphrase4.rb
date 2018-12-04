def part_one
    input = File.read('input/passphrase.in')
    words = Array.new
    valid = 0

    input.each_line do |p|
        words.clear
        valid += 1

        p.split.each do |w|
            if words.include?(w)
                valid -= 1
                break
            end
            words << w
        end
    end

    puts valid
end

def part_two
    input = File.read('input/passphrase.in')
    words = Array.new
    letters = Array.new
    num_valid = 0

    input.each_line do |p|
        words.clear
        letters.clear
        valid = true

        p.split.each do |w|
            crt_letters = Hash.new(0)

            if words.include?(w)
                valid = false
                break
            end

            w.chars do |c|
                crt_letters[c] += 1
            end

            letters.each do |l|
                if l == crt_letters
                    valid = false
                    break
                end
            end

            words << w
            letters << crt_letters
        end

        if valid
            num_valid += 1
        end
    end

    puts num_valid
end

part_two
