#include <stdio.h>
#include <stdlib.h>

typedef struct {
    int start;
    int end;
} Range;

int cmp(const void *a, const void *b) {
    Range *r1 = (Range*)a;
    Range *r2 = (Range*)b;
    return r1->start - r2->start;
}

int mergeRanges(Range ranges[], int n, Range merged[]) {

    qsort(ranges, n, sizeof(Range), cmp);

    int j = 0;
    merged[j] = ranges[0];

    for (int i = 1; i < n; i++) {

        if (ranges[i].start <= merged[j].end) {

            if (ranges[i].end > merged[j].end)
                merged[j].end = ranges[i].end;

        } else {
            j++;
            merged[j] = ranges[i];
        }
    }

    return j + 1;
}

int main() {

    FILE *f = fopen("inputs.txt", "r");

    Range ranges[100000];
    int numbers[100000];

    int rcount = 0, ncount = 0;
    int readingRanges = 1;

    char line[256];

    while (fgets(line, sizeof(line), f)) {

        if (line[0] == '\n') {
            readingRanges = 0;
            continue;
        }

        if (readingRanges) {
            int s, e;
            sscanf(line, "%d-%d", &s, &e);
            ranges[rcount++] = (Range){s,e};
        } else {
            int n;
            sscanf(line, "%d", &n);
            numbers[ncount++] = n;
        }
    }

    // PART 1
    int count = 0;

    for (int i = 0; i < ncount; i++) {
        for (int j = 0; j < rcount; j++) {
            if (numbers[i] >= ranges[j].start && numbers[i] <= ranges[j].end) {
                count++;
                break;
            }
        }
    }

    printf("Part1: %d\n", count);

    // PART 2
    Range merged[100000];

    int mcount = mergeRanges(ranges, rcount, merged);

    int total = 0;

    for (int i = 0; i < mcount; i++) {
        total += merged[i].end - merged[i].start + 1;
    }

    printf("Part2: %d\n", total);

    fclose(f);
}
