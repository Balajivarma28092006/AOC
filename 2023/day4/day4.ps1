$cards = @()

Get-Content "test.txt" | ForEach-Object {
    if($_ -match '^Card\s+(\d+):\s*(.*?)\s\|\s*(.*)$')
    {
        $cardNum = [int]$matches[1]

        $winning = $matches[2] -split '\s+' | ForEach-Object {
            [int]$_
        }

        $mine = $matches[3] -split '\s+' | ForEach-Object {
            [int]$_
        }

        $cards += [PSCustomObject]@{
            Card = $cardNum
            Winning = $winning
            Mine = $mine
        }
    }
}

# the ans is forEach card ans += 2^(matches - 1)
$asnwer = 0

foreach($card in $cards)
{
    $matches = 0
    foreach($num in $card.Mine)
    {
        if($num -in $card.Winning)
        {
            $matches++
        }
    }

    if($matches -eq 0)
    {
        continue
    }
    $answer += [System.Math]::Pow(2, $matches - 1)
}

Write-Host "Part1 :" $answer
