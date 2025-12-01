#include <stdio.h>
#include <string.h>
#include <stdbool.h>

int main()
{
  FILE *fptr = fopen("input.txt", "r");

  char lr;
  int value;

  int zeroCount = 0;
  int dialValue = 50;

  while (fscanf(fptr, " %c%d", &lr, &value) == 2) {
    bool negative = lr == 'L';

    for (int i = 0; i < value; i++)
    {
      if (negative)
      {
        dialValue -= 1;
      } else {
        dialValue += 1;
      }

      if (dialValue % 100 == 0) zeroCount++;
    }
  }

  fclose(fptr);

  printf("ZeroCount: %d", zeroCount);

  return 0;
}

