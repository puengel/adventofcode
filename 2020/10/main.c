#include <stdio.h>
#include <stdlib.h>

int compare(const void *a, const void *b);
int countLines(char *filename);
long long countWays(const int *array, const int startIdx, const int size, long long *seen);

int main()
{

    FILE *fp;
    char *line = NULL;
    size_t len = 0;
    ssize_t read;
    int nLines = countLines("input.txt");

    fp = fopen("input.txt", "r");
    if (fp == NULL)
        exit(EXIT_FAILURE);

    int idx = 0;
    int numbers[nLines];

    while ((read = getline(&line, &len, fp)) != -1)
    {
        numbers[idx++] = atoi(line);
    }

    qsort(numbers, nLines, sizeof(int), compare);

    int before = 0;
    int diff1 = 0;
    int diff3 = 0;
    for (int i = 0; i < nLines; i++)
    {
        if (numbers[i] - before == 1)
            diff1++;
        if (numbers[i] - before == 3)
            diff3++;
        before = numbers[i];
    }

    // last step always 3
    diff3++;

    printf("Part 1: %d\n", (diff1 * diff3));

    long long seen[numbers[nLines - 1]];
    for (int i; i < numbers[nLines - 1]; i++)
    {
        seen[i] = 0;
    }

    int n2[nLines + 1];
    n2[0] = 0;
    for (int i = 0; i < nLines; i++)
    {
        n2[i + 1] = numbers[i];
    }
    long long result = countWays(n2, 0, (nLines + 1), seen);

    printf("Part 2: %lld\n", result);

    fclose(fp);
    if (line)
        free(line);
    exit(EXIT_SUCCESS);
}

long long countWays(const int *array, const int startIdx, const int size, long long *seen)
{
    // last index
    if (startIdx == size - 1)
    {
        seen[startIdx] = 1;
        return 1;
    }

    // calculated already
    if (seen[startIdx] > 0)
    {
        return seen[startIdx];
    }

    long long result = 0;

    for (int i = startIdx + 1; i < size; i++)
    {
        if (array[i] - array[startIdx] <= 3)
        {
            result += countWays(array, i, size, seen);
        }
        else
        {
            break;
        }
    }

    seen[startIdx] = result;

    return result;
}

int compare(const void *a, const void *b)
{
    int int_a = *((int *)a);
    int int_b = *((int *)b);

    // an easy expression for comparing
    return (int_a > int_b) - (int_a < int_b);
}

int countLines(char *filename)
{
    FILE *fp = fopen(filename, "r");
    int ch = 0;
    int lines = 1;

    if (fp == NULL)
        return 0;

    while (!feof(fp))
    {
        ch = fgetc(fp);
        if (ch == '\n')
        {
            lines++;
        }
    }

    fclose(fp);
    return lines;
}