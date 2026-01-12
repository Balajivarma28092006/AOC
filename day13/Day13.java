
import java.io.BufferedReader;
import java.io.FileReader;
import java.util.ArrayList;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Day13 {

    static long solve(List<String> input, boolean  part2){
        long totalCost = 0;
        Pattern numPattern = Pattern.compile("-?\\d+");
        for(int i = 0; i < input.size(); i += 3){
            String block = input.get(i) + input.get(i+1) + input.get(i+2);
            Matcher m = numPattern.matcher(block);
            long[] v = new long[6];
            int idx = 0;
            while(m.find()) v[idx++] = Long.parseLong(m.group());
            long Ax = v[0], Ay = v[1]; // X values
            long Bx = v[2], By = v[3]; // y values
            long Px = v[4], Py = v[5]; // prices

            if (part2){
                 Px += 10_000_000_000_000L;
                Py += 10_000_000_000_000L;
            }

            long D = Ax * By - Ay * Bx; // using crammers rule Determinant
            if (D==0)continue;

            long aNum = Px * By - Py * Bx; // y det
            long bNum = Ax * Py - Px * Ay ; // x det

            if(aNum % D != 0 || bNum % D != 0) continue;

            long a = aNum/D;
            long b = bNum/D;

            if(a < 0 || b < 0) continue;

            totalCost += 3 * a + b;
        }
        return totalCost;
    }

    public static void main(String[] args) throws Exception{
        List<String> input;
        try (BufferedReader br = new BufferedReader(new FileReader("day13_inputs.txt"))) {
            input = new ArrayList<>();
            String line;
            while((line = br.readLine()) != null){
                if(!line.isEmpty()) input.add(line);
            }
        }
        System.out.println("Part 1: " + solve(input, false));
        System.out.println("Part 2: " + solve(input, true));
    }
}
