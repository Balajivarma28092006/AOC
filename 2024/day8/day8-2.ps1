$inputfile = "day8_inputs.txt"

# Read grid
$grid = Get-Content $inputfile
$H = $grid.Count
$W = $grid[0].Length

# Group antennas by frequency
$groups = @{}
for($r = 0; $r -lt $H; $r++){
    for($c = 0; $c -lt $W; $c++){
        $ch = $grid[$r][$c]
        if ($ch -ne '.'){
            if(-not $groups.ContainsKey($ch)){
                $groups[$ch] = New-Object System.Collections.ArrayList
            }
            $null = $groups[$ch].Add([Tuple]::Create($r, $c))
        }
    }
}

# Check if inside grid
function InGrid($r, $c, $H, $W) {
    return ($r -ge 0 -and $c -ge 0 -and $r -lt $H -and $c -lt $W)
}

# gcd helper
function gcd($a, $b){
    while($b -ne 0){
        $t = $b
        $b = $a % $b
        $a = $t
    }
    return [math]::Abs($a)
}

# Part 2 antinode set
$antinodes2 = New-Object System.Collections.Generic.HashSet[string]

foreach($freq in $groups.Keys){
    $points = $groups[$freq]

    # All pairs of antennas
    for ($i = 0; $i -lt $points.Count; $i++) {
        for ($j = $i+1; $j -lt $points.Count; $j++) {

            $p1 = $points[$i]
            $p2 = $points[$j]

            $r1, $c1 = $p1.Item1, $p1.Item2
            $r2, $c2 = $p2.Item1, $p2.Item2

            $dr = $r2 - $r1
            $dc = $c2 - $c1

            # Compute reduced step vector
            $g = gcd $dr $dc
            $stepR = $dr / $g
            $stepC = $dc / $g

            # Add both antennas (Part 2 always includes them)
            $antinodes2.Add("$r1,$c1") | Out-Null
            $antinodes2.Add("$r2,$c2") | Out-Null

            # Walk forward from p2
            $rr = $r2
            $cc = $c2
            while ($true) {
                $rr += $stepR
                $cc += $stepC
                if (-not (InGrid $rr $cc $H $W)) { break }
                $antinodes2.Add("$rr,$cc") | Out-Null
            }

            # Walk backward from p1
            $rr = $r1
            $cc = $c1
            while ($true) {
                $rr -= $stepR
                $cc -= $stepC
                if (-not (InGrid $rr $cc $H $W)) { break }
                $antinodes2.Add("$rr,$cc") | Out-Null
            }
        }
    }
}

Write-Host "Part 2: $($antinodes2.Count)"
$groups
