function solve_part1(input)
    disk = Int[]
    file_id = 0

    # Expand disk
    for (i, c) in enumerate(input)
        size = parse(Int, c)
        if isodd(i)   # file
            append!(disk, fill(file_id, size))
            file_id += 1
        else          # free space
            append!(disk, fill(-1, size))  # -1 = free
        end
    end

    # Two-pointer compaction
    l = 1
    r = length(disk)

    while l < r
        while l < r && disk[l] != -1
            l += 1
        end
        while l < r && disk[r] == -1
            r -= 1
        end

        if l < r
            disk[l], disk[r] = disk[r], disk[l]
            l += 1
            r -= 1
        end
    end

    # Checksum
    checksum = 0
    for i in 1:length(disk)
        if disk[i] != -1
            checksum += (i - 1) * disk[i]
        end
    end

    return checksum
end


function solve_part2(input)
    files = []   # (id, start, length)
    free = []    # (start, length)

    pos = 0
    file_id = 0

    # Parse input
    for (i, c) in enumerate(input)
        size = parse(Int, c)
        if isodd(i)
            push!(files, (file_id, pos, size))
            file_id += 1
        else
            push!(free, (pos, size))
        end
        pos += size
    end

    # Move files right -> left
    for i in length(files):-1:1
        id, fstart, flen = files[i]

        for j in 1:length(free)
            fpos, flen_free = free[j]

            if fpos < fstart && flen_free >= flen
                # move file
                files[i] = (id, fpos, flen)

                # update free segment
                if flen_free == flen
                    deleteat!(free, j)
                else
                    free[j] = (fpos + flen, flen_free - flen)
                end
                break
            end
        end
    end

    # Checksum
    checksum = 0
    for (id, start, len) in files
        for i in start:(start + len - 1)
            checksum += i * id
        end
    end

    return checksum
end


input = strip(read("day9_inputs.txt", String))
println(solve_part1(input))
println(solve_part2(input))
