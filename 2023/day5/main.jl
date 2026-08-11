function parse_input(filename)
    content = read(filename, String)
    sections = split(strip(content), "\n\n")

    # parse the seed
    seed_line = sections[1]
    seeds = [parse(Int, m.match) for m in eachmatch(r"\d+", seed_line)]

    # parse the other 7 mappings
    all_maps = Vector{Vector{Tuple{Int, Int, Int}}}()

    for section in sections[2:end]
        lines = split(section, "\n")
        map_ranges = Tuple{Int, Int, Int}[]

        # skip the first line
        for line in lines[2:end]
            if isempty(strip(line)) continue end
            dest, src, len = map(x -> parse(Int, x), split(line))
            push!(map_ranges, (dest, src, len))
        end
        push!(all_maps, map_ranges)
    end
    return seeds, all_maps
end

function solve_part1(filename)
    seeds, all_maps = parse_input(filename)
    final_seed = Int[]

    for seed in seeds
        curr_seed = seed 
        for layer in all_maps
            matched = false
            for (dest, src, len) in layer
                if src <= curr_seed <= src + len 
                    curr_seed = dest + (curr_seed - src) # spawn after dest
                    matched = true
                    break
                end
            end
        end
        push!(final_seed, curr_seed)
    end
    return minimum(final_seed)
end

println("part 1: ", solve_part1("inputs.txt"))

function solve_part2(filename)
    seed, _ = solve_part1(filename)
    testthis = Int[]
    

    for i in []
end