using Printf

function max_number_from_digits(s::String, k::Int)
    digits = collect(s)
    n = length(digits)

    result = Char[]
    start = 1

    for i in 1:k
        last = n - ( k - i )

        max_digit = '0'
        max_index = start

        for j in start:last
            if digits[j] > max_digit
                max_digit = digits[j]
                max_index = j
            end
        end

        push!(result, max_digit)
        start = max_index + 1
    end

    return join(result)
end

function main()
    input = read("inputs.txt", String)
    lines = split(chomp(input), '\n')

    total = 0
    for line in lines
        val = parse(Int, max_number_from_digits(String(line), 2))
        total += val
    end
    @printf("Part1: %.2f\n", total)
end

main()
