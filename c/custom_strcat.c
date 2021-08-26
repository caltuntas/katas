#include <stdio.h>
#include <string.h>
char* custom_strcat(char* dest, const char* src){
   char* ptr = dest + strlen(dest);
   while(*src != '\0') {
      *ptr++ = *src++;
   }

   *ptr = '\0';
   return dest;
}
int main() {
   char str1[100] = "Hello", str2[] = " world";

   custom_strcat(str1, str2);

   puts(str1);
   puts(str2);

   return 0;
}
