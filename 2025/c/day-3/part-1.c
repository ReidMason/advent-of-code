#include <stdio.h>
#include <string.h>
#include "stringlist.h"

int getLineBatterySum(char lineValue[]) {
  int lineLength = strlen(lineValue);
  int max = 0;

  for (int i = 0; i < lineLength-1; i++) {
    char character = lineValue[i];
    int digit = character - '0';

    for (int j = i+1; j < lineLength; j++) {
      char character2 = lineValue[j];
      int digit2 = character2 - '0';

      int sum = digit * 10 + digit2;
      if (sum > max) {
        max = sum;
      }
    }
  }

  return max;
}

int solve(char input[]) {
  StringList *line = split(input, "\n");

  int total = 0;

  while (line) {
    char *lineValue = line->value;
    total += getLineBatterySum(lineValue);

    line = line->next;
  }

  return total;
}

void test() {
  char input[] = "987654321111111\n811111111111119\n234234234234278\n818181911112111";

  int actual = solve(input);

  int expected = 357;

  printf("Test %s\n Actual: %d\n", expected != actual ? "Failed" : "Passed", actual);
}


int main() {
  // test();

  FILE *fptr = fopen("input.txt", "r");

  char line[100];

  // Get file size
  fseek(fptr, 0, SEEK_END);
  long fileSize = ftell(fptr);
  fseek(fptr, 0, SEEK_SET);

  // Allocate memory and read entire file
  char *input = malloc(fileSize + 1);
  fread(input, 1, fileSize, fptr);
  input[fileSize] = '\0';

  fclose(fptr);

  int total = solve(input);

  printf("Answer: %d\n", total);

  return 0;
}

