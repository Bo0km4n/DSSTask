#include <stdlib.h>
#include <string.h>
#include "util.h"
#include <string.h>
#include <stdlib.h>
#include "descriptor.h"


char *newnstr(const unsigned char *bytes, const size_t len) {
    char *ret = malloc(sizeof(unsigned char) * (len + 1));
    memcpy(ret, bytes, len);
    memset(&ret[len], '\0', 1);
    return ret;
}

char *newstr(const char *str) {
    size_t len = strlen(str) + 1;
    char *ret = malloc(sizeof(char) * len);
    strcpy(ret, str);
    return ret;
}

struct bytes_t bytes_from_char(unsigned char from) {
    struct bytes_t bytes;
    bytes.len = 1;
    bytes.head = malloc(sizeof(char) * bytes.len);
    bytes.head[0] = from;
    return bytes;
}

struct bytes_t bytes_from_short(unsigned short from) {
    struct bytes_t bytes;
    bytes.len = 2;
    bytes.head = malloc(sizeof(char) * bytes.len);
    for (size_t i = 1; i <= bytes.len; i++) {
        bytes.head[bytes.len - i] = (unsigned char)from;
        from >>= BYTE;
    }
    return bytes;
}

struct bytes_t bytes_from_int(unsigned int from) {
    struct bytes_t bytes;
    bytes.len = 4;
    bytes.head = malloc(sizeof(char) * bytes.len);
    for (size_t i = 1; i <= bytes.len; i++) {
        bytes.head[bytes.len - i] = (unsigned char)from;
        from >>= BYTE;
    }
    return bytes;
}

struct bytes_t bytes_from_long(unsigned long from) {
    struct bytes_t bytes;
    bytes.len = 8;
    bytes.head = malloc(sizeof(char) * bytes.len);
    for (size_t i = 1; i <= bytes.len; i++) {
        bytes.head[bytes.len - i] = (unsigned char)from;
        from >>= BYTE;
    }
    return bytes;
}

struct bytes_t bytes_from_string(const char *str) {
    struct bytes_t bytes;
    size_t slen = strlen(str);
    bytes.len = slen + UTF_HEADER_SIZE;
    bytes.head = malloc(sizeof(char) * bytes.len);
    memcpy(&bytes.head[UTF_HEADER_SIZE], str, slen);
    for (size_t i = 1; i <= UTF_HEADER_SIZE; i++) {
        bytes.head[UTF_HEADER_SIZE - i] = (unsigned char)slen;
        slen >>= BYTE;
    }
    return bytes;
}

void hexdump(const char *desc, const unsigned char *pc, const size_t len) {
  int i;
  unsigned char buff[17];
  if (desc != NULL)
    printf ("%s:\n", desc);
  if (len == 0) {
    printf("  ZERO LENGTH\n");
    return;
  }
  // if (len < 0) {
  //   printf("  NEGATIVE LENGTH: %i\n",len);
  //   return;
  // }
  for (i = 0; i < len; i++) {
    if ((i % 16) == 0) {
      if (i != 0)
        printf ("  %s\n", buff);
      printf ("  %04x ", i);
    }
    printf (" %02x", pc[i]);
    if ((pc[i] < 0x20) || (pc[i] > 0x7e))
      buff[i % 16] = '.';
    else
      buff[i % 16] = pc[i];
    buff[(i % 16) + 1] = '\0';
  }
  while ((i % 16) != 0) {
    printf ("   ");
    i++;
  }
  printf ("  %s\n", buff);
}
