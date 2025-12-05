#include <stdlib.h>

struct StringList
{
  struct StringList *next;
  char value[];
};

typedef struct StringList StringList;

StringList* newStringList(char value[])
{
  StringList *list = malloc(sizeof(StringList) + strlen(value) + 1);
  list->next = NULL;
  strcpy(list->value, value);
  return list;
}

void setValue(StringList *stringList, char value[])
{
  strcpy(stringList->value, value);
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
