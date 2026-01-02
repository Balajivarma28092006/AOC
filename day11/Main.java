/*

* rule1: 0 -> 1
* rule2: if the numbers are even number of digits abcd -> ab | cd
* rule3: if the number digits are odd then x -> x * 2024
* rule4: repeat it for every blink

 */

import java.io.BufferedReader;
import java.io.FileReader;
import java.util.HashMap;
import java.util.Map;

public class Main {

    static final int BLINKS_PART1 = 25;
    static final int BLINKS_PART2 = 75;

    public static void main(String[] args) {
        String filename = "day11.txt";
        try (BufferedReader br = new BufferedReader(new FileReader(filename))) {
            String line = br.readLine();
            String[] parts = line.split("\\s+");

            // so instead of storing the stone values we gonna store the number of time a stone value appeared
            Map<Long, Long> freq = new HashMap<>();
            for (String p : parts) {
                long val = Long.parseLong(p);
                freq.put(val, freq.getOrDefault(val, 0L) + 1);
            }

            /*

            * consider this example as
              value  count
              -----  -----
              125      1
              17       1

            * converted to this
              value  count
              253000   1 as 125 (has odd digits then 125 -> 125*2024)
              1        1 and 17 -> 1 | 7
              7        1

             */


            //generate blinks
            for(int i = 0; i < BLINKS_PART2; i++) { //toggle part 1-2 for other part answers
                freq = blinks(freq);
            }

            long answer = 0;
            for(long count : freq.values()) answer+=count;
            System.out.println(answer);
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }

    private static Map<Long, Long> blinks (Map<Long, Long> freq) {
        Map<Long, Long> next = new HashMap<>();

        for (Map.Entry<Long, Long> entry : freq.entrySet()) {
            long stone = entry.getKey();
            long count = entry.getValue();

            //Rule 1
            if (stone == 0){
                add(next, 1, count);
                continue;
            }

            String s = Long.toString(stone);

            //Rule 2: even digits -> split
            if (s.length() % 2 == 0) {
                int mid = s.length() / 2;
                long left = Long.parseLong(s.substring(0, mid));
                long right = Long.parseLong(s.substring(mid));

                add(next, left, count);
                add(next, right, count);
            }

            //Rule 3: odd digits -> x*2024
            else {
                add(next, stone*2024, count);
            }
        }
        return next;
    }

    private static void add(Map<Long, Long> map, long key, long value) {
        map.put(key, map.getOrDefault(key, 0L) + value);
    }
}