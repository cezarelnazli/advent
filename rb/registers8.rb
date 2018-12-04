input = File.read('input/registers.in')
all_max = 0
registers = Hash.new(0)

input.each_line do |l|
    reg, op, val, _, c_reg, c_op, c_val = l.split
    val = val.to_i
    c_val = c_val.to_i
    if registers[c_reg].send c_op, c_val
        if op == 'dec'
            registers[reg] -= val
        else
            registers[reg] += val
        end

        if registers[reg] > all_max
            all_max = registers[reg]
        end
    end
end

p all_max
p registers.values.max
