dirs = [(-1, 0), (1, 0), (0, -1), (0, 1)]

function read_input(filename)
    lines = readlines(filename)
    rows = length(lines)
    cols = length(lines[1])

    grid = zeros(Int, rows, cols)
    for i in 1:rows
        for j in 1:cols
            grid[i, j] = lines[i][j] - '0'
        end
    end
    return grid
end

function solve_part1(grid, rows, cols)
    total = 0
    for i in 1:rows
        for j in 1:cols
            if grid[i, j] == 0
                reached = Set{Tuple{Int,Int}}()
                dfsPart1(grid, rows, cols, i, j, reached)
                total += length(reached)
            end
        end
    end
    return total
end

function dfsPart1(grid, rows, cols, r, c, reached)
    if grid[r, c] == 9
        push!(reached, (r, c))
        return
    end

    for (dr, dc) in dirs
        nr = r + dr
        nc = c + dc

        if 1 <= nr <= rows && 1 <= nc <= cols && grid[nr, nc] == grid[r, c] + 1
            dfsPart1(grid, rows, cols, nr, nc, reached)
        end
    end
end

function solve_part2(grid, rows, cols)
    total = 0
    for i in 1:rows
        for j in 1:cols
            if grid[i, j] == 0
                total += dfsPart2(grid, rows, cols, i, j)
            end
        end
    end
    return total
end


function dfsPart2(grid, rows, cols, i, j)
    path = 0
    if grid[i, j] == 9
        return 1
    end

    for (dr, dc) in dirs
        nr = i + dr
        nc = j + dc

       if 1 <= nr <= rows && 1 <= nc <= cols && grid[nr, nc] == grid[i, j] + 1
            path += dfsPart2(grid, rows, cols, nr, nc)
        end
    end
    return path
end





grid = read_input("day10_inputs.txt")
rows, cols = size(grid)

println("Part 1 = ", solve_part1(grid, rows, cols))
println("Part 2 = ", solve_part2(grid, rows, cols))
