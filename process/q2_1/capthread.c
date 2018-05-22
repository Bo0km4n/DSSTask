#include "capthread.h"

int main(void) {
    pid_t pid;

    int pp[2];      // pipe
    char buf[256];

    printf("$ type something... >>>");
    fgets(buf, 256, stdin);

    pthread_t pthread;
    pthread_create( &pthread, NULL, &capitalize, buf);
    pthread_join(pthread, NULL);
    printf("%s", buf);
}


char * capitalize_str(char *c) {
    char *p;
    for(p=c; *p; p++) {
        *p = (char)toupper((int)*p);
    }
    return c;
}

// capitalize
void capitalize(void *str) {
    char *buf = (char *)str;
    buf = capitalize_str(buf);
}
