import java.io.BufferedReader;
import java.io.FileReader;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;

public class Day15 {

    static char[][] grid;
    static String moves;
    static int robotRow;
    static int robotCol;

    static Map<Character, int[]> directions = Map.of(
        '^', new int[]{-1, 0},
        'v', new int[]{1, 0},
        '<', new int[]{0, -1},
        '>', new int[]{0, 1}
    );

    static void readGrid() throws Exception {
        List<String> lines = new ArrayList<>();

        BufferedReader br = new BufferedReader(new FileReader("grid.txt"));

        String line;
        while ((line = br.readLine()) != null) {
            if (!line.isEmpty()) {
                lines.add(line);
            }
        }

        br.close();

        grid = new char[lines.size()][];

        for (int i = 0; i < lines.size(); i++) {
            grid[i] = lines.get(i).toCharArray();
        }
    }

    static void readMoves() throws Exception {
        BufferedReader br = new BufferedReader(new FileReader("directions.txt"));

        StringBuilder sb = new StringBuilder();

        String line;
        while ((line = br.readLine()) != null) {
            sb.append(line);
        }

        br.close();

        moves = sb.toString().replaceAll("[^<>^v]", "");
    }

    static void findRobot() {
        for (int r = 0; r < grid.length; r++) {
            for (int c = 0; c < grid[r].length; c++) {
                if (grid[r][c] == '@') {
                    robotRow = r;
                    robotCol = c;
                    return;
                }
            }
        }
    }

    static void move(char dir) {
        int[] d = directions.get(dir);

        int dr = d[0];
        int dc = d[1];

        int nr = robotRow + dr;
        int nc = robotCol + dc;

        if (grid[nr][nc] == '#') {
            return;
        }

        if (grid[nr][nc] == '.') {
            grid[robotRow][robotCol] = '.';
            grid[nr][nc] = '@';

            robotRow = nr;
            robotCol = nc;
            return;
        }

        if (grid[nr][nc] == 'O') {
            int br = nr;
            int bc = nc;

            while (grid[br][bc] == 'O') {
                br += dr;
                bc += dc;
            }

            if (grid[br][bc] == '.') {
                grid[br][bc] = 'O';

                grid[nr][nc] = '@';
                grid[robotRow][robotCol] = '.';

                robotRow = nr;
                robotCol = nc;
            }
        }
    }

    static int gps() {
        int total = 0;

        for (int r = 0; r < grid.length; r++) {
            for (int c = 0; c < grid[r].length; c++) {
                if (grid[r][c] == 'O') {
                    total += 100 * r + c;
                }
            }
        }

        return total;
    }

    public static void main(String[] args) throws Exception {
        readGrid();
        readMoves();
        findRobot();

        for (char move : moves.toCharArray()) {
            move(move);
        }

        System.out.println(gps());
    }
}