#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>
#include "stringlist.h"

int getIntDigits(int value) {
  int digits = 0;

  while (value != 0) {
    digits++;
    value = value / 10;
  }

  return digits;
}

StringList* split(char text[], char delimeter[]) {
  StringList *outputRoot = NULL;
  StringList *outputHead = NULL;
  char* token = strtok(text, delimeter);

  while (token != NULL) {
    StringList *newNode = newStringList(token);

    if (outputRoot == NULL)
    {
      outputRoot = newNode;
      outputHead = newNode;
    } else {
      outputHead->next = newNode;
      outputHead = newNode;
    }

    token = strtok(NULL, delimeter);
  }

  return outputRoot;
}

long long int solve(char input[])
{
  StringList *list = split(input, ",");

  long long int total = 0;

  while (list) {
    // printf("%s\n", list->value);

    StringList *numbers = split(list->value, "-");
    long long int start = atoll(&numbers->value[0]);
    long long int end = atoll(&numbers->next->value[0]);

    // printf(" start: %d end: %d\n", start, end);
    for (long long int i = start; i <= end; i++)
    {
      int digits = getIntDigits(i);
      char numstr[digits];
      sprintf(numstr,"%lld",i);

      // printf("  checking: %s\n", numstr); 
      int length = strlen(numstr);
      for (int j = 0; j < length/2; j++)
      {
        char section[j+1];
        for (int k = 0; k <= j; k++)
        {
          section[k] = numstr[k];
        }
        // printf("   chunk: %s\n", section);

        int lengthOfSection = strlen(section);
        
        int repetitions = 2; // length / lengthOfSection;
        // printf("    reps: %d sectionLen: %d\n", repetitions, lengthOfSection);
        char tesselated[digits];
        for (int k = 0; k < repetitions; k++)
        {
          for (int l = 0; l < lengthOfSection; l++)
          {
            // printf("    k: %d los: %d l: %d idx: %d\n", k, lengthOfSection, l, k * lengthOfSection + l);
            tesselated[k * lengthOfSection + l] = section[l];
          }
        }

        // printf("    tesselated: %s\n", tesselated);

        long long int tesselsatedInt = atoll(tesselated);
        if (tesselsatedInt == i)
        {
          total += i;
        }
      }
    }

    list = list->next;
  }

  return total;
}

void test() {
  char input[] = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";

  long long int expected = 1227775554;

  long long int actual = solve(input);

  printf("Test %s\n actual: %lld\n", expected != actual ? "Failed" : "Passed", actual);
}


int main()
{
  FILE *fptr = fopen("input.txt", "r");

  char input[1028];
  while (fgets(input, 1028, fptr)) {
    printf("%s", input);
  }

  fclose(fptr);

  long long int answer = solve(input);

  printf("Answer: %lld\n", answer);

  // Wrong answers
  // 571808556 - Too low
  // 12586854255 - Winner winner chicken dinner

  return 0;
}

