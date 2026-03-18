#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define MAXC 200
#define MAXR 200

char grid[MAXR][MAXC];
bool visited[MAXR][MAXC] = { false };

int dr[4] = { -1, 1, 0, 0 };
int dc[4] = { 0, 0, -1, 1 };

int flood_area, flood_perimeter;
char target;

void dfs(int r, int c, int R, int C){
    visited[r][c] = true;
    flood_area++;

    for(int i = 0; i < 4; i++){
        int nr = r + dr[i];
        int nc = c + dc[i];

        if(nr < 0 || nr >= R || nc < 0 || nc >= C || grid[nr][nc] != target){
            flood_perimeter++;
        } else if (!visited[nr][nc]) {
            dfs(nr, nc, R, C);
        }
    }
}

int main() {
    FILE *fp;

    fp = fopen("day12_inputs.txt", "r");
    if (fp == NULL){
        perror("Error opening file!");
        exit(1);
    }
    
    char line[256];
    int R = 0, C = 0;

    while ( fgets(line, sizeof(line), fp)) {
        line[strcspn(line, "\n")] = 0;
        C = strlen(line);

        for(int i = 0; i < C; i++){
            grid[R][i] = line[i];
        }
        R++;
    }
    
    long long total_price = 0;
    for(int r = 0; r < R; r++){
        for(int c = 0; c < C; c++){
            if(!visited[r][c]){
                target = grid[r][c];
                flood_area = 0;
                flood_perimeter = 0;
                dfs(r, c, R, C);
                total_price += (long long)flood_area*flood_perimeter;
            }
        }
    }
    
    printf("%lld\n", total_price);
    fclose(fp);
    return 0;
}
