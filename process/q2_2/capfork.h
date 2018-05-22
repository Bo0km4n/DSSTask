#ifndef CAPFORK_H
#define CAPFORK_H

#include <sys/types.h>
#include <unistd.h>
#include <stdio.h>
#include <string.h>
#include <ctype.h>
#include <stdlib.h>


int write_to_pipe(int pp[2], char *m);
int read_from_pipe(int pp[2]);
void capitalize(void *str);
#endif
