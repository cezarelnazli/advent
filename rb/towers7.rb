require 'set'
require 'scanf'

$input = File.read('input/towers.in')

def part_one
    candidates = Set.new
    sustained = Set.new

    $input.each_line do |l|
        parts = l.split(" -> ")

        candidates << parts[0].split[0].strip

        if parts.length == 2
            parts[1].split(", ").each do |p|
                sustained << p.strip
            end
        end
    end

    (candidates - sustained).first
end

class Node
    attr_accessor :weight, :sum, :children
end

def make_sum(root)
    root.sum = root.weight

    root.children.each do |c|
        make_sum(c)
        root.sum += c.sum
    end
end

def check_sum(root)
    queue = Array.new
    head = 0

    queue << root

    loop do
        if head >= queue.length
            break
        end

        node = queue[head]
        head += 1

        if node.children.length > 0
            benchmark = 0

            if node.children.first.sum == node.children[1].sum
                benchmark = node.children.first.sum
            else
                if node.children.first.sum == node.children[2].sum
                    benchmark = node.children.first.sum
                else
                    benchmark = node.children[1].sum
                end
            end

            node.children.each do |c|
                queue << c

                if c.sum != benchmark
                    w, b = check_sum(c)
                    if b
                        return c.weight - (c.sum - benchmark), false
                    else
                        return w, false
                    end
                end
            end
        end
    end

    return root.weight, true
end

def part_two
    nodes = Hash.new

    $input.each_line do |l|
        parts = l.split(' -> ')

        program, weight = parts.first.scanf('%s (%d)')

        node = nodes[program]
        if node != nil
            node.weight = weight
        else
            nodes[program] = Node.new
            nodes[program].weight = weight
            nodes[program].sum = 0
            nodes[program].children = Array.new
            node = nodes[program]
        end

        if parts.length == 2
            parts[1].strip.split(', ').each do |p|
                n = nodes[p]
                if n == nil
                    nodes[p] = Node.new
                    nodes[p].weight = 0
                    nodes[p].sum = 0
                    nodes[p].children = Array.new
                end

                nodes[program].children << nodes[p]
            end
        end
    end

    root = nodes[part_one()]
    make_sum(root)
    r, _ = check_sum(root)
    p r
end

part_two
