"""
Advent of Code 2025 - Day 4
Look at every @ in the grid.
Count how many neighboring cells also contain @.
If that count is less than 4, add it to the answer.
"""

directons = [
    (-1, -1), (-1, 0), (-1, 1),
    (0, -1),           (0, 1),
    (1, -1), (1, 0), (1, 1)
]

function solve_part1(grid)
    count = 0
    rows = length(grid)
    cols = length(grid[1])

    for i in 1:rows
        for j in 1:cols
            if grid[i][j] == '@'
                neighbor_count = 0
                for (dx, dy) in directons
                    x = i + dx
                    y = j + dy
                    if 1 <= x <= rows && 1 <= y <= cols && grid[x][y] == '@'
                        neighbor_count += 1
                    end
                end
                if neighbor_count < 4
                    count += 1
                end
            end
        end
    end
    
    return count
end

function solve_part2(grid)
    total_removed = 0
    rows = length(grid)
    cols = length(grid[1])

    while true
        to_remove = Tuple{Int, Int}[]
        for i in 1:rows
            for j in 1:cols
                if grid[i][j] == '@'
                    neighbor_count = 0
                    for (dx, dy) in directons
                        x = i + dx
                        y = j + dy
                        if 1 <= x <= rows && 1 <= y <= cols && grid[x][y] == '@'
                            neighbor_count += 1
                        end
                    end
                    if neighbor_count < 4
                        push!(to_remove, (i, j))
                    end
                end
            end
        end

        if isempty(to_remove)
            break
        end

        for (i, j) in to_remove
            grid[i][j] = '.'
        end
        total_removed += length(to_remove)
    end
    return total_removed
end


function main()
    lines = readlines("inputs.txt")
    grid = [collect(line) for line in lines]
    println("Part 1: ", solve_part1(grid))

    
    grid = [collect(line) for line in lines]
    println("Part 2: ", solve_part2(grid))
end

main()