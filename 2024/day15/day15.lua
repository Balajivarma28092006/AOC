local Warehouse = {}
Warehouse.__index = Warehouse

local directions = {
    ["^"] = {-1, 0},
    ["v"] = {1, 0},
    ["<"] = {0, -1},
    [">"] = {0, 1}
}

function Warehouse:new()
    local obj = setmetatable({}, self)

    obj.grid = obj:read_grid()
    obj.moves = obj:read_moves()
    obj.robot_row, obj.robot_col = obj:find_robot()

    return obj
end

function Warehouse:read_grid()
    local file = assert(io.open("grid.txt", "r"))

    local grid = {}

    for line in file:lines() do
        if line ~= "" then
            local row = {}

            for i = 1, #line do
                row[i] = line:sub(i, i)
            end

            table.insert(grid, row)
        end
    end

    file:close()

    return grid
end

function Warehouse:read_moves()
    local file = assert(io.open("directions.txt", "r"))

    local content = file:read("*all")
    file:close()

    return content:gsub("[^<>^v]", "")
end

function Warehouse:find_robot()
    for r = 1, #self.grid do
        for c = 1, #self.grid[r] do
            if self.grid[r][c] == "@" then
                return r, c
            end
        end
    end

    error("Robot not found")
end

function Warehouse:move(dir)
    local dr = directions[dir][1]
    local dc = directions[dir][2]

    local nr = self.robot_row + dr
    local nc = self.robot_col + dc

    local cell = self.grid[nr][nc]

    if cell == "#" then
        return
    end

    if cell == "." then
        self.grid[self.robot_row][self.robot_col] = "."
        self.grid[nr][nc] = "@"

        self.robot_row = nr
        self.robot_col = nc
        return
    end

    if cell == "O" then
        local br = nr
        local bc = nc

        while self.grid[br][bc] == "O" do
            br = br + dr
            bc = bc + dc
        end

        if self.grid[br][bc] == "." then
            self.grid[br][bc] = "O"
            self.grid[nr][nc] = "@"
            self.grid[self.robot_row][self.robot_col] = "."

            self.robot_row = nr
            self.robot_col = nc
        end
    end
end

function Warehouse:gps_sum()
    local total = 0

    for r = 1, #self.grid do
        for c = 1, #self.grid[r] do
            if self.grid[r][c] == "O" then
                total = total + 100 * (r - 1) + (c - 1)
            end
        end
    end

    return total
end

local warehouse = Warehouse:new()

for move in warehouse.moves:gmatch(".") do
    warehouse:move(move)
end

print(warehouse:gps_sum())