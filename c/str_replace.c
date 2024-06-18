#include <string.h>
#include <stdio.h>
#include <stdlib.h>

char* strreplace(char *str, char* search, char* replacement)
{
  char *newstr = (char*)malloc(strlen(str));

  const char* tmp = str;
  const char* from = str;
  int previndex = 0;
  int replacement_len = strlen(replacement);
  int search_len = strlen(search);
  while((tmp = strstr(tmp,search)))
  {
    int index = (int)(tmp - str);

    strncat(newstr,from,index - previndex);
    strncat(newstr,replacement,replacement_len);
    from = from + index - previndex + search_len;
    tmp++;
    previndex = index + search_len;
  }

  if (previndex < strlen(str))
    strncat(newstr,from,strlen(str)-previndex);

  return newstr;
}

int main()
{
  char *name = "mynameis@testtest@com";

  const char* tmp = name;

  char* rep = strreplace(name,"@","---");
  printf("result is %s\n",rep);

  return 0;
}
