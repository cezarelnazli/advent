input = File.read('input/maze.in')

instructions = input.each_line.map { |s| s.to_i }

steps = 0
cpos = 0

loop do
    if cpos < 0 || cpos >= instructions.length
        break
    end

    if instructions[cpos] >= 3
        instructions[cpos] -= 1
        cpos += instructions[cpos] + 1
    else
        instructions[cpos] += 1
        cpos += instructions[cpos] - 1
    end
    steps += 1
end

puts steps
