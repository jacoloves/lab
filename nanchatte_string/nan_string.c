#include <stdio.h>
#include <string.h>
#include <stddef.h>

typedef struct {
  char* target_str;
} nan_string;

//char* nan_strProto(char* target_str, char* copy_str);

int main()
{
  nan_string test_str = {"sfaadfad"};


  printf("%s\n", test_str.target_str);
  return 0;
}

char* nan_strProto(char* target_str, char* copy_str)
{
  memcpy(&target_str, &copy_str, sizeof(copy_str));

  return target_str;
}
