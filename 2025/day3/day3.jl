function best_number(line, k)
    digits = [parse(Int, c) for c in line if isdigit(c)]
    drop = length(digits) - k
    stack = Int[]

# <<<<<<< HEAD
    for d in digits
        while drop > 0 && !isempty(stack) && stack[end] < d
            pop!(stack)
            drop -= 1
        end
        push!(stack, d)
    end

    stack = stack[1:min(k, length(stack))]

    val = 0
    for d in stack
        val = val * 10 + d
    end
    return val
end
# =======
# Optimized Part 1: O(n) using a Monotonic Stack with a length limit
# To get the largest number of length K, we remove (length(s) - k) digits.
function part1(s::AbstractString, k::Int)
    n = length(s)
    if k >= n return string(s) end
    if k <= 0 return "" end

    to_remove = n - k
    stack = Char[]

    for c in s
        # While we still have digits to remove AND the current digit 
        # is bigger than the last one we kept:
        while to_remove > 0 && !isempty(stack) && stack[end] < c
            pop!(stack)
            to_remove -= 1
        end
        push!(stack, c)
    end

    # If we still need to remove digits (e.g., input was "9876"), 
    # remove them from the end.
    return join(stack[1:k])
end

# Part 2: The "Joltage" Rule
# Based on your examples, you want the largest possible subsequence.
# This is usually the Monotonic Stack without a length limit.
function part2(s::AbstractString)
    if isempty(s) return "0" end
    stack = Char[]

    for c in s
        while !isempty(stack) && stack[end] < c
            pop!(stack)
        end
        push!(stack, c)
    end
    
    result = join(stack)
    return isempty(result) ? "0" : result
end

function main()
    filename = "inputs.txt"
    if !isfile(filename)
        println("Error: File not found.")
        return
    end

    lines = filter(!isempty, split(read(filename, String), '\n'))
    
    total1 = BigInt(0)
    total2 = BigInt(0)

    for line in lines
        clean_line = strip(line)
        if isempty(clean_line) continue end

        # PART 1: 
        # Double check your AOC instructions: 
        # Does Part 1 always ask for length 12, or is k variable?
        p1_str = part1(clean_line, 12)
        total1 += parse(BigInt, p1_str)

        # PART 2:
        p2_str = part2(clean_line)
        total2 += parse(BigInt, p2_str)
    end

    @printf("Part1: %s\n", string(total1))
    @printf("Part2: %s\n", string(total2))
end

lines = readlines("inputs.txt")

main()