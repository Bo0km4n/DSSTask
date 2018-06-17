#ifndef NET_H
#define NET_H

#define PORT 18888
#define HOST "http://localhost"
#include <stdio.h>
typedef struct
{
    char* memory;
    size_t size;
} Memory;

void fetch_struct(Memory *buf);
void store_struct(void);

#endif