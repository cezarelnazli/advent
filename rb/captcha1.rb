def step_one
    number = File.read('input/captcha.in')
    sum = 0

    if number[0] == number[number.length - 1]
        sum += number[0].to_i
    end

    i = 0
    loop do
        if number[i] == number[i + 1]
            sum += number[i].to_i
        end

        i += 1
        if i == number.length - 1
            break
        end
    end

    puts sum
end

def step_two
    number = File.read('input/captcha.in')
    sum = 0
    step = number.length / 2

    i = 0
    loop do
        if number[i] == number[i + step]
            sum += number[i].to_i
        end

        i += 1
        if i == step
            break
        end
    end

    puts 2 * sum
end
