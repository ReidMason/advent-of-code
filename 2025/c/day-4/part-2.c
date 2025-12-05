#include <stdio.h>
#include <string.h>
#include <math.h>
#include "stringlist.h"

int getMaxIndex(char line[], int start, int end) {
  int max = 0;
  int idx = 0;
  for (int i = start; i < end; i++) {
    char character = line[i];
    int digit = character - '0';
    if (digit > max) {
      max = digit;
      idx = i;
    }
  }

  return idx;
}

long long int getLineBatterySum(char line[]) {
  int lineLength = strlen(line);

  int digits = 12;
  int prevMaxIdx = -1;
  long long int total = 0;

  for (int i = digits - 1; i >= 0; i--) {
    int startIdx = prevMaxIdx + 1; 
    prevMaxIdx = getMaxIndex(line, startIdx, lineLength - i);
    char max = line[prevMaxIdx];
    int maxDigit = max - '0';

    total += maxDigit * pow(10, i);
  }

  return total;
}

long long int solve(char input[]) {
  StringList *line = split(input, "\n");

  long long int total = 0;

  while (line) {
    char *lineValue = line->value;
    total += getLineBatterySum(lineValue);

    line = line->next;
  }

  return total;
}

void test() {
  char input[] = "987654321111111\n811111111111119\n234234234234278\n818181911112111";

  long long int actual = solve(input);

  long long int expected = 3121910778619;

  printf("Test %s\n Actual: %lld\n", expected != actual ? "Failed" : "Passed", actual);
}


int main() {
  // test();
  // return 0;

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

  long long int total = solve(input);

  printf("Answer: %lld\n", total);

  return 0;
}

