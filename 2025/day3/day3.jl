function best_number(line, k)
    digits = [parse(Int, c) for c in line if isdigit(c)]
    drop = length(digits) - k
    stack = Int[]

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

lines = readlines("inputs.txt")

p1 = sum(best_number(line, 2) for line in lines)
p2 = sum(best_number(line, 12) for line in lines)

println("Part 1: ", p1)
println("Part 2: ", p2)