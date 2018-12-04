input = File.read('input/memory.in')

steps = 0
banks = input.split.map { |x| x.to_i }
banks_cache = Array.new { Array.new }

loop do
    banks_cache << Array.new(banks)

    imax = 0
    banks.each_index do |i|
        if banks[i] > banks[imax]
            imax = i
        end
    end

    emax = banks[imax]
    banks[imax] = 0
    banks.each_index do |i|
        banks[i] += emax / banks.length
        if (i - imax > 0 && i - imax <= emax % banks.length) ||
                (i - imax <= 0 && i <= emax % banks.length - (banks.length - imax))
            banks[i] += 1
        end
    end

    steps += 1
    banks_cache.each_index do |i|
        if banks_cache[i] == banks
            puts banks_cache.length - i
            puts steps
            exit
        end
    end
end
