#include "net.h"
#include <stdio.h>
#include <stdlib.h>
int main(void) {
    Memory *buf = (Memory *)malloc(sizeof(Memory));
    buf->memory = NULL;
    buf->size = 0;
    
    fetch_struct(buf);

    printf("fetched data: \n");
    for (int i = 0; i < buf->size; i++) {
        printf("0x%02x ", buf->memory[i] & 0x000000FF);
    }
    printf("\n");
}