#ifndef CAPFORK_H
#define CAPFORK_H

#include <sys/types.h>
#include <unistd.h>
#include <stdio.h>
#include <string.h>
#include <ctype.h>
#include <stdlib.h>

typedef struct {
    int *pp;
    char *buf;
    char *type;
} p_message;

int write_to_pipe(p_message *m);
void read_from_pipe(p_message *m);
void capitalize(void *str);
#endif
