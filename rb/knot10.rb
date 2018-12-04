$input = File.read('input/knot.in')

LISTSIZE = 256

def one_round(lengths, cpos, skip, list)
    lengths.each do |l|
        aux_list = Array.new(LISTSIZE)

        limit = cpos + l
        limit_first = 0
        if limit > LISTSIZE
            limit_first = limit % LISTSIZE
            limit = LISTSIZE
        end

        j = 0
        i = cpos
        loop do
            if i >= limit
                break
            end

            aux_list[j] = list[i]
            j += 1
            i += 1
        end

        i = 0
        loop do
            if i >= limit_first
                break
            end

            aux_list[j] = list[i]
            j += 1
            i += 1
        end

        j -= 1
        i = cpos
        loop do
            if i >= limit
                break
            end

            list[i] = aux_list[j]
            j -= 1
            i += 1
        end

        i = 0
        loop do
            if i >= limit_first
                break
            end

            list[i] = aux_list[j]
            j -= 1
            i += 1
        end

        cpos = (cpos + l + skip) % LISTSIZE
        skip += 1
    end

    return lengths, cpos, skip, list
end

def part_one
    list = Array.new(LISTSIZE) { |i| i }
    lengths = Array.new
    cpos = 0
    skip = 0

    $input.split(',').each do |l|
        lengths << l.to_i
    end

    one_round lengths, cpos, skip, list

    p list[0] * list[1]
end

def part_two
    list = Array.new(LISTSIZE) { |i| i }
    lengths = Array.new
    cpos = 0
    skip = 0

    $input.strip.each_byte do |b|
        lengths << b
    end

    lengths << 17
    lengths << 31
    lengths << 73
    lengths << 47
    lengths << 23

    64.times do
        lengths, cpos, skip, list = one_round lengths, cpos, skip, list
    end

    16.times do |i|
        n = list[i<<4]
        15.times do |j|
            n ^= list[(i << 4) + (j + 1)]
        end

        printf '%0.2x', n
    end

    puts ''
end

part_one
part_two
