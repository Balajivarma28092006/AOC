
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class Day10 {

    static int[][] grid;
    static int rows, cols;

    //directions
    static int[] dr = {-1, 1, 0, 0};
    static int[] dc = {0, 0, -1, 1};

    public static void main(String[] args) {
        Path filepath = Path.of("day10_inputs.txt");

        try {
        List<String> lines = Files.readAllLines(filepath);
        // String[] data = lines.toArray(String[]::new);
        rows = lines.size();
        cols = lines.get(0).length();
        grid = new int[rows][cols];

        for (int i = 0; i < rows; i++) {
            for (int j = 0; j < cols; j++) {
                grid[i][j] = lines.get(i).charAt(j) - '0';
            }
        }

        System.out.println("Part 1: " + SolvePart1());
        System.out.println("Part 2: " + SolvePart2());

        } catch (IOException e) {
            System.err.println("Error: opening the file " + e.getMessage());
        }
    }

    static int SolvePart1() {
        int total = 0;
        for (int i = 0; i < rows; i++) {
            for (int j = 0; j < cols; j++) {
                if (grid[i][j] == 0) {
                    Set<String> reachedNines = new HashSet<>();
                    dfsPart1(i, j, reachedNines);
                    // System.out.println(reachedNines);
                    total += reachedNines.size();
                }
            }
        }
        return total;
    }

    static void dfsPart1(int r, int c, Set<String> reachedNines) {
        if (grid[r][c] == 9) {
            reachedNines.add(r + "," + c);
        }

        for (int d = 0; d < 4; d++) {
            int nr = r + dr[d];
            int nc = c + dc[d];

            if (nc >= 0 && nc < cols && nr >= 0 && nr < rows && grid[nr][nc] == grid[r][c] + 1) {
                dfsPart1(nr, nc, reachedNines);
            }
        }
    }

    static int SolvePart2() {
        int total = 0;
        for (int i = 0; i < rows; i++) {
            for (int j = 0; j < cols; j++) {
                if (grid[i][j] ==0 ){
                    total += dfsPart2(i, j);
                }
            }
        }
        return total;
    }

    static int dfsPart2(int r, int c) {
         if (grid[r][c] == 9) {
            return 1;
        }
        int paths = 0;
        for (int d = 0; d < 4; d++) {
            int nr = r + dr[d];
            int nc = c + dc[d];

            if (nc >= 0 && nc < cols && nr >= 0 && nr < rows && grid[nr][nc] == grid[r][c] + 1) {
                paths += dfsPart2(nr, nc);
            }
        }
        return paths;
    }
}
