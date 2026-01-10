#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX 200

int H = 0, W = 0;
char grid[MAX][MAX];
int visited[MAX][MAX];

int dx[4] = {0, 0, -1, 1};
int dy[4] = {-1, 1, 0, 0};

/* boundary[dir][y][x] = 1 if boundary edge exists */
int boundary[4][MAX][MAX];

void dfs(int y, int x, char plant, int *area) {
    visited[y][x] = 1;
    (*area)++;

    for (int d = 0; d < 4; d++) {
        int ny = y + dy[d];
        int nx = x + dx[d];

        if (ny < 0 || ny >= H || nx < 0 || nx >= W ||
            grid[ny][nx] != plant) {
            boundary[d][y][x] = 1;
        } else if (!visited[ny][nx]) {
            dfs(ny, nx, plant, area);
        }
    }
}

void dfs_side(int d, int y, int x) {
    boundary[d][y][x] = 0;

    for (int i = 0; i < 4; i++) {
        int ny = y + dy[i];
        int nx = x + dx[i];

        if (ny >= 0 && ny < H && nx >= 0 && nx < W &&
            boundary[d][ny][nx]) {
            dfs_side(d, ny, nx);
        }
    }
}

int main(int argc, char *argv[]) {
    if (argc != 2) {
        fprintf(stderr, "Usage: %s <input_file>\n", argv[0]);
        return 1;
    }

    FILE *fp = fopen(argv[1], "r");
    if (!fp) {
        perror("Error opening file");
        return 1;
    }

    /* read grid from file */
    while (fgets(grid[H], MAX, fp)) {
        grid[H][strcspn(grid[H], "\n")] = 0;
        W = strlen(grid[H]);
        H++;
    }
    fclose(fp);

    long long answer = 0;

    for (int y = 0; y < H; y++) {
        for (int x = 0; x < W; x++) {
            if (visited[y][x]) continue;

            memset(boundary, 0, sizeof(boundary));

            int area = 0;
            dfs(y, x, grid[y][x], &area);

            int sides = 0;
            for (int d = 0; d < 4; d++) {
                for (int i = 0; i < H; i++) {
                    for (int j = 0; j < W; j++) {
                        if (boundary[d][i][j]) {
                            sides++;
                            dfs_side(d, i, j);
                        }
                    }
                }
            }

            answer += (long long)area * sides;
        }
    }

    printf("%lld\n", answer);
    return 0;
}
