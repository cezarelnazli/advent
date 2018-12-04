input = File.read('input/stream.in')

level = 0
sum = 0
garbage = 0
ignore_next = false
in_garbage = false

input.each_char do |c|
    if ignore_next
        ignore_next = false
        next
    end

    if in_garbage
        case c
        when '!'
            ignore_next = true
        when '>'
            in_garbage = false
        else
            garbage += 1
        end

        next
    end

    case c
    when '!'
        ignore_next = true
    when '<'
        in_garbage = true
    when '{'
       level += 1
       sum += level
    when '}'
        level -= 1
    end
end

p sum
p garbage
