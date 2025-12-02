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
