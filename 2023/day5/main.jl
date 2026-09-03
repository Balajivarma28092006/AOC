function parse_input(filename)
    content = read(filename, String)
    sections = split(strip(content), "\n\n")

    # parse the seed
    seed_line = sections[1]
    seeds = [parse(Int, m.match) for m in eachmatch(r"\d+", seed_line)]

    # parse the other 7 mappings
    all_maps = Vector{Vector{Tuple{Int,Int,Int}}}()

    for section in sections[2:end]
        lines = split(section, "\n")
        map_ranges = Tuple{Int,Int,Int}[]

        # skip the first line
        for line in lines[2:end]
            if isempty(strip(line))
                continue
            end
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
    # the main idea was brute force which is worse so i took llm support
    # shame on myself ... its third time i went to llm since i started AOC damn.. means three days worth of knowledge
    # without any struggle
    seeds, all_maps = parse_input(filename)

    # convert the seeds to (start, end) ranges 
    current_ranges = Tuple{Int,Int}[]
    for i in 1:2:length(seeds)
        push!(current_ranges, (seeds[i], seeds[i+1]))
    end

    for layer in all_maps
        next_ranges = Tuple{Int,Int}[]

        # do bfs
        while !isempty(current_ranges)
            r_start, r_len = pop!(current_ranges)
            r_end = r_start + r_len - 1
            matched = false

            for (dest, src, len) in layer
                src_end = src + len - 1

                #check for any intersection between the seeds and map range 
                overlap_start = max(r_start, src)
                overlap_end = min(r_end, src_end)
                maximum
                if overlap_start <= overlap_end
                    matched = true

                    overlap_len = overlap_end - overlap_start + 1
                    offset = dest - src
                    push!(next_ranges, (overlap_start + offset, overlap_len))

                    if r_start < overlap_start
                        push!(current_ranges, (r_start, overlap_start - r_start))
                    end

                    if r_end > overlap_end
                        push!(current_ranges, (overlap_end + 1, r_end - overlap_end))
                    end

                    break
                end
            end
            if !matched
                push!(next_ranges, (r_start, r_len))
            end
        end
        current_ranges = next_ranges
    end
    return minimum(r[1] for r in current_ranges)
end

println(solve_part2("inputs.txt"))
