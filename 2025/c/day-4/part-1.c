#include <stdio.h>
#include <string.h>
#include <stdbool.h>
#include "stringlist.h"

int adjustOffset(int offset, int currentIndex, int lineLength, int linePos) {
  int linePosOffset = linePos + offset;
  if (linePosOffset > lineLength) {
    return offset + 1;
  } else if (linePosOffset < 0) {
    return offset - 1;
  }

  return offset;
}

bool isIndexValid(int index, int inputLength) {
  return index > 0 && index <= inputLength;
}

int solve(char input[]) {
  int inputLength = strlen(input);
  StringList *line = split(input, "\n");

  int total = 0;
  int lineLength = strlen(line->value);
  lineLength--;

  int rows = inputLength/lineLength;
  printf("rows: %d", rows);

  for (int i = 0; i < inputLength; i++) {
    char value = input[i];

    if (value != '@') {
      continue;
    }

    int linePos = (i % lineLength) + 1;

    int offsets[8] = { -lineLength-1, -lineLength, -lineLength+1, -1, 1, lineLength-1, lineLength, lineLength+1 };

    int neighbours = 0;
    for (int j = 0; j < 8; j++) {
      int offset = offsets[j];
      offset = adjustOffset(offset, i, lineLength, linePos);
      int index = i + offset;
      char item = 'n';
      if (isIndexValid(index, inputLength)) {
        item = input[index];
      }

      // printf("%d %c\n", index, item);

      if (item == '@') {
        neighbours++;
      }
    }

    bool valid = neighbours < 4;
    // printf("LinePos: %d %s\n", linePos, valid ? "valid" : "invalid");

    if (valid) {
      printf("%d\n", linePos);
      total++;
    }
  }

  // printf("%d", total);
  return total;
}

void test() {
  char input[] = "..@@.@@@@.\n@@@.@.@.@@\n@@@@@.@.@@\n@.@@@@..@.\n@@.@@@@.@@\n.@@@@@@@.@\n.@.@.@.@@@\n@.@@@.@@@@\n.@@@@@@@@.\n@.@.@@@.@.";

  int actual = solve(input);

  int expected = 13;

  printf("Test %s\n Actual: %d\n", expected != actual ? "Failed" : "Passed", actual);
}


int main() {
  test();
  return 0;

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

