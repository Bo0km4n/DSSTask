#ifndef UTIL_H
#define UTIL_H

#include <stdio.h>

struct bytes_t {
  unsigned char *head;
  size_t len;
};

struct bytes_t bytes_from_char(unsigned char from);
struct bytes_t bytes_from_short(unsigned short from);
struct bytes_t bytes_from_int(unsigned int from);
struct bytes_t bytes_from_long(unsigned long from);
struct bytes_t bytes_from_string(const char *str);

char *newnstr(const unsigned char *bytes, const size_t len);
char *newstr(const char *str);

void hexdump(const char *desc, const unsigned char *pc, const size_t len);
#endif