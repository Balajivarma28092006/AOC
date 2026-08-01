cards = Dict{Int, NamedTuple{(:winning, :mine), Tuple{Vector{Int}, Vector{Int}}}}()

open("inputs.txt", "r") do file
	for line in eachline(file)
		if isempty(strip(line)) continue end

		card_part, number_part = split(line, ":")
		card_num = parse(Int, match(r"\d+", card_part).match)

		winning_str, mine_str = split(number_part, "|")
		winning_nums = [parse(Int, m.match) for m in eachmatch(r"\d+", winning_str)]
		mine_nums = [parse(Int, m.match) for m in eachmatch(r"\d+", mine_str)]

		cards[card_num] = (winning = winning_nums, mine = mine_nums)
	end
end

match_count = Dict{Int, Int}()
for (id, card) in cards
	match_count[id] = length(intersect(card.winning, card.mine))
end

copies = fill(1, length(match_count))
n = length(match_count)
for card in 1:n 
	wins = match_count[card]
	for next in card + 1: card + wins
		copies[next] += copies[card]
	end
end

println(sum(copies))
