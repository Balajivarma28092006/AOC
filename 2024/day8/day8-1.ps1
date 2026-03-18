$inputfile = "day8_inputs.txt"

$grid = Get-Content $inputfile
$H = $grid.Count
$W = $grid[0].Length

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

function InGrid($r, $c, $H, $W) {
    return ($r -ge 0 -and $c -ge 0 -and $r -lt $H -and $c -lt $W)
}

function gcd($a, $b){
    while($b -ne 0){
        $t = $b
        $b = $a % $b
        $a = $t
    }
    return [math]::Abs($a)
}

$antinodes1 = New-Object System.Collections.Generic.HashSet[string]

foreach($freq in $groups.Keys){
    $points = $groups[$freq]

    for ($i = 0; $i -lt $points.Count; $i++) {
        for ($j = $i+1; $j -lt $points.Count; $j++) {

            $p1 = $points[$i]
            $p2 = $points[$j]

            $r1, $c1 = $p1.Item1, $p1.Item2
            $r2, $c2 = $p2.Item1, $p2.Item2

            $dr = $r2 - $r1
            $dc = $c2 - $c1

            # Part 1 antinodes
            $a1r = $r1 - $dr
            $a1c = $c1 - $dc

            $a2r = $r2 + $dr
            $a2c = $c2 + $dc

            if (InGrid $a1r $a1c $H $W) {
                $antinodes1.Add("$a1r,$a1c") | Out-Null
            }
            if (InGrid $a2r $a2c $H $W) {
                $antinodes1.Add("$a2r,$a2c") | Out-Null
            }
        }
    }
}

Write-Host "Part 1: $($antinodes1.Count)"
