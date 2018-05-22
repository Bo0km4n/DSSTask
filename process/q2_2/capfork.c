#include "capfork.h"

int main(void) {
    pid_t pid;

    int pp[2];      // pipe
    char buf[256];

    pipe(pp);       // make pipe
    pid = fork();   // make child process
    
    if (pid == 0) {
        printf("$ type your message >>> ");
        fgets(buf, 256, stdin);
        write_to_pipe(pp, buf);
    } else {
        read_from_pipe(pp);
    }
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

int write_to_pipe(int pp[2], char *m) {
    close(pp[0]);
    capitalize(m);
    printf("write message : %s", m);
    write(pp[1], m, strlen(m) + 1);
    return close(pp[1]);
}

int read_from_pipe(int pp[2]) {
    close(pp[1]);
    char buf[256];
    read(pp[0], buf, 256);
    printf("read message from pipe: %s", buf);
    return close(pp[0]);
}