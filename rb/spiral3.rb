# Solve part 2 of Advent of Code day 3 problem. Part one was solved more easily
# using the interpreter, and did not require building and walking the matrix.
# One comment on Reddit says that there exists an OEIS entry for this problem as
# well, so no code is required here either.
INPUT = 312051

# Size of matrix such that it contains the input (312051) and the solution.
# It is currently "big enough", but I should find a way to compute the proper
# size.
#
DIMENSION = 101
m = Array.new(DIMENSION) { Array.new(DIMENSION, 0) }

# When computing M[i][j] as the sum of all adjacent numbers including diagonals,
# these are the position differences inside the matrix.
#
SUM_DIRECTIONS = [
    [-1, -1],   # NW
    [-1, +0],   # N
    [-1, +1],   # NE
    [+0, -1],   # W
    [+0, +1],   # E
    [+1, -1],   # SW
    [+1, +0],   # S
    [+1, +1]    # SE
]

# Begin the walk at the center of the matrix, of coordinates i and j below.
#
i = (DIMENSION - 1) / 2
j = (DIMENSION - 1) / 2

# We are on level 1 (the innermost square, consisting of just one number).
#
lvl = 1

# At each iteration, We are placing the pos'th number (initially, pos = 1),
# which is n (initially, n = 1).
#
pos = 1
n = 1

# We are not walking in any direction currently, but during the first iteration,
# we immediately move to level 2, which updates the walk direction (we're going
# up).
#
walk = nil

# How many steps can we take in the current direction (steps_limit)
# and how many steps we've already taken (steps_taken).
#
steps_limit = 0
steps_taken = 1

loop do
    # Place new value in the matrix. If we found what we're looking for (first n
    # that is greater than input), print the answer and stop the loop.
    #
    m[i][j] = n
    if n > INPUT
        puts n
        break
    end

    # If we're on a matrix corner (the current position follows the formula
    # below), we get on the next level (outer square), move to the right, the
    # next direction is up, can take 2 more steps than before and we've already
    # taken one (moving to the outer square).
    if (2 * lvl - 1)**2 == pos
        lvl += 1
        j += 1
        walk = :up
        steps_limit += 2
        steps_taken = 1
    # Otherwise, update i and j according to the current direction. If we've
    # reached the steps limit in the current direction, update the direction.
    else
        if walk == :up
            i -= 1
        elsif walk == :down
            i += 1
        elsif walk == :left
            j -= 1
        else
            j += 1
        end

        steps_taken += 1
        if steps_taken == steps_limit
            steps_taken = 0
            if walk == :up
                walk = :left
            elsif walk == :down
                walk = :right
            elsif walk == :left
                walk = :down
            else
                walk = :up
            end
        end
    end

    # Next iteration computes next number.
    #
    pos += 1

    # Update n by inspecting all adjacent numbers using the directions array.
    #
    n = 0
    SUM_DIRECTIONS.each do |d|
        n += m[i + d[0]][j + d[1]]
    end
end
