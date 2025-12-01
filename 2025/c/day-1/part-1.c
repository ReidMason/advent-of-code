#include <stdio.h>
#include <string.h>

int main()
{
  FILE *fptr = fopen("input.txt", "r");

  char lr;
  int value;

  int zeroCount = 0;
  int dialValue = 50;

  while (fscanf(fptr, " %c%d", &lr, &value) == 2) {
    if (lr == 'L') {
      value = -value;
    }

    dialValue += value;
    if (dialValue % 100 == 0) zeroCount++;
  }

  fclose(fptr);

  printf("ZeroCount: %d", zeroCount);

  return 0;
}

