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

println(parse_input("test.txt"))